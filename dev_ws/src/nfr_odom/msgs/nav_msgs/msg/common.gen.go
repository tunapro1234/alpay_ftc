// Code generated by rclgo-gen. DO NOT EDIT.

package nav_msgs_msg

/*
#cgo LDFLAGS: "-L/home/ros/dev_ws/install/mecanum_control_go/lib" "-Wl,-rpath=/home/ros/dev_ws/install/mecanum_control_go/lib"
#cgo LDFLAGS: "-L/home/ros/dev_ws/install/cleaner_navigation/lib" "-Wl,-rpath=/home/ros/dev_ws/install/cleaner_navigation/lib"
#cgo LDFLAGS: "-L/home/ros/dev_ws/install/cleaner_bringup/lib" "-Wl,-rpath=/home/ros/dev_ws/install/cleaner_bringup/lib"
#cgo LDFLAGS: "-L/home/ros/dev_ws/install/cleaner_description/lib" "-Wl,-rpath=/home/ros/dev_ws/install/cleaner_description/lib"
#cgo LDFLAGS: "-L/opt/ros/humble/lib" "-Wl,-rpath=/opt/ros/humble/lib"

#cgo LDFLAGS: -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lnav_msgs__rosidl_typesupport_c -lnav_msgs__rosidl_generator_c
#cgo LDFLAGS: -lbuiltin_interfaces__rosidl_typesupport_c -lbuiltin_interfaces__rosidl_generator_c
#cgo LDFLAGS: -lgeometry_msgs__rosidl_typesupport_c -lgeometry_msgs__rosidl_generator_c
#cgo LDFLAGS: -lstd_msgs__rosidl_typesupport_c -lstd_msgs__rosidl_generator_c

#cgo CFLAGS: "-I/home/ros/dev_ws/install/mecanum_control_go/include/action_msgs"
#cgo CFLAGS: "-I/home/ros/dev_ws/install/mecanum_control_go/include/builtin_interfaces"
#cgo CFLAGS: "-I/home/ros/dev_ws/install/mecanum_control_go/include/example_interfaces"
#cgo CFLAGS: "-I/home/ros/dev_ws/install/mecanum_control_go/include/geometry_msgs"
#cgo CFLAGS: "-I/home/ros/dev_ws/install/mecanum_control_go/include/rosidl_runtime_c"
#cgo CFLAGS: "-I/home/ros/dev_ws/install/mecanum_control_go/include/rosidl_typesupport_interface"
#cgo CFLAGS: "-I/home/ros/dev_ws/install/mecanum_control_go/include/sensor_msgs"
#cgo CFLAGS: "-I/home/ros/dev_ws/install/mecanum_control_go/include/std_msgs"
#cgo CFLAGS: "-I/home/ros/dev_ws/install/mecanum_control_go/include/std_srvs"
#cgo CFLAGS: "-I/home/ros/dev_ws/install/mecanum_control_go/include/test_msgs"
#cgo CFLAGS: "-I/home/ros/dev_ws/install/mecanum_control_go/include/unique_identifier_msgs"
#cgo CFLAGS: "-I/home/ros/dev_ws/install/mecanum_control_go/include/builtin_interfaces"

#cgo CFLAGS: "-I/home/ros/dev_ws/install/mecanum_control_go/include/geometry_msgs"

#cgo CFLAGS: "-I/home/ros/dev_ws/install/mecanum_control_go/include/std_msgs"

#cgo CFLAGS: "-I/home/ros/dev_ws/install/mecanum_control_go/include/nav_msgs"

#cgo CFLAGS: "-I/home/ros/dev_ws/install/cleaner_navigation/include/action_msgs"
#cgo CFLAGS: "-I/home/ros/dev_ws/install/cleaner_navigation/include/builtin_interfaces"
#cgo CFLAGS: "-I/home/ros/dev_ws/install/cleaner_navigation/include/example_interfaces"
#cgo CFLAGS: "-I/home/ros/dev_ws/install/cleaner_navigation/include/geometry_msgs"
#cgo CFLAGS: "-I/home/ros/dev_ws/install/cleaner_navigation/include/rosidl_runtime_c"
#cgo CFLAGS: "-I/home/ros/dev_ws/install/cleaner_navigation/include/rosidl_typesupport_interface"
#cgo CFLAGS: "-I/home/ros/dev_ws/install/cleaner_navigation/include/sensor_msgs"
#cgo CFLAGS: "-I/home/ros/dev_ws/install/cleaner_navigation/include/std_msgs"
#cgo CFLAGS: "-I/home/ros/dev_ws/install/cleaner_navigation/include/std_srvs"
#cgo CFLAGS: "-I/home/ros/dev_ws/install/cleaner_navigation/include/test_msgs"
#cgo CFLAGS: "-I/home/ros/dev_ws/install/cleaner_navigation/include/unique_identifier_msgs"
#cgo CFLAGS: "-I/home/ros/dev_ws/install/cleaner_navigation/include/builtin_interfaces"

#cgo CFLAGS: "-I/home/ros/dev_ws/install/cleaner_navigation/include/geometry_msgs"

#cgo CFLAGS: "-I/home/ros/dev_ws/install/cleaner_navigation/include/std_msgs"

#cgo CFLAGS: "-I/home/ros/dev_ws/install/cleaner_navigation/include/nav_msgs"

#cgo CFLAGS: "-I/home/ros/dev_ws/install/cleaner_bringup/include/action_msgs"
#cgo CFLAGS: "-I/home/ros/dev_ws/install/cleaner_bringup/include/builtin_interfaces"
#cgo CFLAGS: "-I/home/ros/dev_ws/install/cleaner_bringup/include/example_interfaces"
#cgo CFLAGS: "-I/home/ros/dev_ws/install/cleaner_bringup/include/geometry_msgs"
#cgo CFLAGS: "-I/home/ros/dev_ws/install/cleaner_bringup/include/rosidl_runtime_c"
#cgo CFLAGS: "-I/home/ros/dev_ws/install/cleaner_bringup/include/rosidl_typesupport_interface"
#cgo CFLAGS: "-I/home/ros/dev_ws/install/cleaner_bringup/include/sensor_msgs"
#cgo CFLAGS: "-I/home/ros/dev_ws/install/cleaner_bringup/include/std_msgs"
#cgo CFLAGS: "-I/home/ros/dev_ws/install/cleaner_bringup/include/std_srvs"
#cgo CFLAGS: "-I/home/ros/dev_ws/install/cleaner_bringup/include/test_msgs"
#cgo CFLAGS: "-I/home/ros/dev_ws/install/cleaner_bringup/include/unique_identifier_msgs"
#cgo CFLAGS: "-I/home/ros/dev_ws/install/cleaner_bringup/include/builtin_interfaces"

#cgo CFLAGS: "-I/home/ros/dev_ws/install/cleaner_bringup/include/geometry_msgs"

#cgo CFLAGS: "-I/home/ros/dev_ws/install/cleaner_bringup/include/std_msgs"

#cgo CFLAGS: "-I/home/ros/dev_ws/install/cleaner_bringup/include/nav_msgs"

#cgo CFLAGS: "-I/home/ros/dev_ws/install/cleaner_description/include/action_msgs"
#cgo CFLAGS: "-I/home/ros/dev_ws/install/cleaner_description/include/builtin_interfaces"
#cgo CFLAGS: "-I/home/ros/dev_ws/install/cleaner_description/include/example_interfaces"
#cgo CFLAGS: "-I/home/ros/dev_ws/install/cleaner_description/include/geometry_msgs"
#cgo CFLAGS: "-I/home/ros/dev_ws/install/cleaner_description/include/rosidl_runtime_c"
#cgo CFLAGS: "-I/home/ros/dev_ws/install/cleaner_description/include/rosidl_typesupport_interface"
#cgo CFLAGS: "-I/home/ros/dev_ws/install/cleaner_description/include/sensor_msgs"
#cgo CFLAGS: "-I/home/ros/dev_ws/install/cleaner_description/include/std_msgs"
#cgo CFLAGS: "-I/home/ros/dev_ws/install/cleaner_description/include/std_srvs"
#cgo CFLAGS: "-I/home/ros/dev_ws/install/cleaner_description/include/test_msgs"
#cgo CFLAGS: "-I/home/ros/dev_ws/install/cleaner_description/include/unique_identifier_msgs"
#cgo CFLAGS: "-I/home/ros/dev_ws/install/cleaner_description/include/builtin_interfaces"

#cgo CFLAGS: "-I/home/ros/dev_ws/install/cleaner_description/include/geometry_msgs"

#cgo CFLAGS: "-I/home/ros/dev_ws/install/cleaner_description/include/std_msgs"

#cgo CFLAGS: "-I/home/ros/dev_ws/install/cleaner_description/include/nav_msgs"

#cgo CFLAGS: "-I/opt/ros/humble/include/action_msgs"
#cgo CFLAGS: "-I/opt/ros/humble/include/builtin_interfaces"
#cgo CFLAGS: "-I/opt/ros/humble/include/example_interfaces"
#cgo CFLAGS: "-I/opt/ros/humble/include/geometry_msgs"
#cgo CFLAGS: "-I/opt/ros/humble/include/rosidl_runtime_c"
#cgo CFLAGS: "-I/opt/ros/humble/include/rosidl_typesupport_interface"
#cgo CFLAGS: "-I/opt/ros/humble/include/sensor_msgs"
#cgo CFLAGS: "-I/opt/ros/humble/include/std_msgs"
#cgo CFLAGS: "-I/opt/ros/humble/include/std_srvs"
#cgo CFLAGS: "-I/opt/ros/humble/include/test_msgs"
#cgo CFLAGS: "-I/opt/ros/humble/include/unique_identifier_msgs"
#cgo CFLAGS: "-I/opt/ros/humble/include/builtin_interfaces"

#cgo CFLAGS: "-I/opt/ros/humble/include/geometry_msgs"

#cgo CFLAGS: "-I/opt/ros/humble/include/std_msgs"

#cgo CFLAGS: "-I/opt/ros/humble/include/nav_msgs"
*/
import "C"
