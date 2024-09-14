from launch import LaunchDescription
from launch.actions import IncludeLaunchDescription, GroupAction, DeclareLaunchArgument
from launch_ros.actions import PushRosNamespace
from launch.launch_description_sources import PythonLaunchDescriptionSource
from ament_index_python.packages import get_package_share_directory
from launch.substitutions import PathJoinSubstitution, LaunchConfiguration, TextSubstitution
from launch.conditions import IfCondition

def generate_launch_description():
    # Paths to various files and directories
    gazebo_ros_pkg = get_package_share_directory('gazebo_ros')
    alpay_bringup_pkg = get_package_share_directory('alpay_bringup')
    diffbot_description_pkg = get_package_share_directory('diffbot_description')

    # Define arguments
    # use_namespace = LaunchConfiguration('use_namespace')
    # namespace = LaunchConfiguration('namespace')

    # Create the launch description
    return LaunchDescription([
        # Declare launch arguments
        # DeclareLaunchArgument('use_namespace', default_value='false', description='Use a namespace if true'),
        # DeclareLaunchArgument('namespace', default_value='simulation1', description='Namespace to use'),

        # Push namespace if use_namespace is true
        # PushRosNamespace(
        #     condition=IfCondition(use_namespace),
        #     namespace=namespace
        # ),

        # Run Gazebo
        IncludeLaunchDescription(
            PythonLaunchDescriptionSource(
                PathJoinSubstitution([gazebo_ros_pkg, 'launch', 'gazebo.launch.py'])
            ),
            launch_arguments={
                'world': PathJoinSubstitution([alpay_bringup_pkg, 'worlds', 'basic_world.world']),
                'params-file': PathJoinSubstitution([alpay_bringup_pkg, 'param', 'gazebo_params.yaml']),
            }.items(),
        ),

        # Group for robot1 with namespace
        GroupAction([
            IncludeLaunchDescription(
                PythonLaunchDescriptionSource(
                    PathJoinSubstitution([diffbot_description_pkg, 'launch', 'sim.launch.py'])
                ), launch_arguments={
                    # 'use_namespace': 'true',
                    # 'namespace': namespace,
                }.items(),
            )
        ]),
    ])
