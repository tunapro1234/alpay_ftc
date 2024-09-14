from ament_index_python.packages import get_package_share_directory
from launch import LaunchDescription
from launch.actions import IncludeLaunchDescription, TimerAction, DeclareLaunchArgument
from launch.launch_description_sources import PythonLaunchDescriptionSource
from launch.substitutions import Command, LaunchConfiguration, TextSubstitution
from launch.actions import RegisterEventHandler
from launch.event_handlers import OnProcessStart
from launch_ros.actions import Node
from launch.conditions import IfCondition
import os

def generate_launch_description():
    package_name='diffbot_description'

    controller_params_file = os.path.join(get_package_share_directory("diffbot_description"), "param", "diffbot_controllers_sim.yaml")

    # Define arguments
    use_namespace = LaunchConfiguration('use_namespace')
    namespace = LaunchConfiguration('namespace')
    namespace_str = TextSubstitution(text=namespace)

    return LaunchDescription([
        # Declare launch arguments
        DeclareLaunchArgument('use_namespace', default_value='false', description='Use a namespace if true'),
        DeclareLaunchArgument('namespace', default_value='robot1', description='Namespace to use'),

        # Push namespace if use_namespace is true
        PushRosNamespace(
            condition=IfCondition(use_namespace),
            namespace=namespace
        ),

        # Controller Manager Node
        TimerAction(period=0.0, actions=[
            Node(
                package="controller_manager",
                executable="ros2_control_node",
                parameters=[controller_params_file],
                remappings=[
                    (
                        "/controller_manager/robot_description",
                        "/robot_description"
                    ),
                ]
            )
        ]),

        # Differential Drive Controller
        TimerAction(period=3.0, actions=[
            Node(
                package="controller_manager",
                executable="spawner",
                arguments=["diff_cont"],
            )
        ]),

        # Joint State Brodcaster Controller
        TimerAction(period=3.0, actions=[
            Node(
                package="controller_manager",
                executable="spawner",
                arguments=["joint_broad"],
            )
        ]),
    ])
