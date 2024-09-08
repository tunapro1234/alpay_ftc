import os

from ament_index_python.packages import get_package_share_directory


from launch import LaunchDescription
from launch.actions import IncludeLaunchDescription, TimerAction
from launch.launch_description_sources import PythonLaunchDescriptionSource
from launch.substitutions import Command
from launch.actions import RegisterEventHandler
from launch.event_handlers import OnProcessStart

from launch_ros.actions import Node



def generate_launch_description():
    package_name='cleaner_bringup'

    controller_params_file = os.path.join(get_package_share_directory("cleaner_description"), "param", "gultepe_controllers.yaml")

    controller_manager = Node(
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

    delayed_controller_manager = TimerAction(period=0.0, actions=[controller_manager])

    #delayed_controller_manager = TimerAction(period=3.0, actions=[controller_manager])

    diff_drive_spawner = Node(
        package="controller_manager",
        executable="spawner",
        arguments=["diff_cont"],
    )

    delayed_diff_drive_spawner  = TimerAction(period=3.0, actions=[diff_drive_spawner])

    #delayed_diff_drive_spawner = RegisterEventHandler(
    #    event_handler=OnProcessStart(
    #        target_action=controller_manager,
    #        on_start=[diff_drive_spawner],
    #    )
    #)

    joint_broad_spawner = Node(
        package="controller_manager",
        executable="spawner",
        arguments=["joint_broad"],
    )
    
    delayed_joint_broad_spawner  = TimerAction(period=3.0, actions=[joint_broad_spawner])

    #delayed_joint_broad_spawner = RegisterEventHandler(
    #    event_handler=OnProcessStart(
    #        target_action=controller_manager,
    #        on_start=[joint_broad_spawner],
    #    )
    #)


    # Code for delaying a node (I haven't tested how effective it is)
    # 
    # First add the below lines to imports
    # from launch.actions import RegisterEventHandler
    # from launch.event_handlers import OnProcessExit
    #
    # Then add the following below the current diff_drive_spawner
    # delayed_diff_drive_spawner = RegisterEventHandler(
    #     event_handler=OnProcessExit(
    #         target_action=spawn_entity,
    #         on_exit=[diff_drive_spawner],
    #     )
    # )
    #
    # Replace the diff_drive_spawner in the final return with delayed_diff_drive_spawner



    # Launch them all!
    return LaunchDescription([
        delayed_controller_manager,
        delayed_diff_drive_spawner,
        delayed_joint_broad_spawner
    ])
