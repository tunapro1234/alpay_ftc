from launch import LaunchDescription
from launch_ros.actions import Node, PushRosNamespace
from launch.substitutions import Command, PathJoinSubstitution, TextSubstitution
from ament_index_python.packages import get_package_share_directory
from launch.actions import DeclareLaunchArgument, IncludeLaunchDescription
from launch.launch_description_sources import PythonLaunchDescriptionSource
from launch.substitutions import LaunchConfiguration, PythonExpression
from launch.conditions import IfCondition


def generate_launch_description():
    # Get the package directories
    diffbot_description_pkg = get_package_share_directory('diffbot_description')
    controller_manager_pkg = get_package_share_directory('controller_manager')

    # Define arguments
    use_namespace = LaunchConfiguration('use_namespace')
    namespace = LaunchConfiguration('namespace')

    # URDF and RViz configuration
    urdf_file = PathJoinSubstitution([diffbot_description_pkg, 'urdf', 'diffbot.urdf.xacro'])
    rviz_config_file = PathJoinSubstitution([diffbot_description_pkg, 'rviz', 'alpay_default.rviz'])

    # Default values 
    robot_description_topic = "/robot_description"
    entity_name = "diffbot"

    robot_description_topic = "/robot_description"
    robot_description_topic = PythonExpression([
        '"/robot_description" if "', use_namespace, '" == "false" else "/', namespace, '/robot_description"'
    ])
    
    entity_name = PythonExpression([
        '"diffbot" if "', use_namespace, '" == "false" else "', namespace, '_diffbot"'
    ])

    return LaunchDescription([
        # Declare launch arguments
        DeclareLaunchArgument('use_namespace', default_value='false', description='Use a namespace if true'),
        DeclareLaunchArgument('namespace', default_value='robot1', description='Namespace to use'),

        # Push namespace if use_namespace is true
        PushRosNamespace(
            condition=IfCondition(use_namespace),
            namespace=namespace
        ),

        # Robot State Publisher
        Node(
            package='robot_state_publisher',
            executable='robot_state_publisher',
            parameters=[{'robot_description': Command(['xacro ', urdf_file])}],
        ),

        # Gazebo Entity Spawner
        Node(
            package='gazebo_ros',
            executable='spawn_entity.py',
            arguments=['-topic', robot_description_topic, '-entity', "diffbot", '-x', '0', '-y', '0', '-z', '0.01'],
        ),

        # RViz
        # Node(
        #     package='rviz2',
        #     executable='rviz2',
        #     output='screen',
        #     arguments=['-d', rviz_config_file],
        # ),

        # Include the controller_manager_sim.launch.py and pass the namespace
        # IncludeLaunchDescription(
        #     PythonLaunchDescriptionSource(PathJoinSubstitution([controller_manager_pkg, 'launch', 'controller_manager_sim.launch.py'])),
        #     launch_arguments={'use_namespace': use_namespace, 'namespace': namespace}.items()
        # )
    ])
