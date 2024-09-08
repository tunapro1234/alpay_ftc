cmake_minimum_required(VERSION 3.5)
project(mecanum_control_go)

# Find the necessary ROS 2 packages
find_package(ament_cmake REQUIRED)

# Set the source directory and output binary
set(SOURCE_DIR ${CMAKE_CURRENT_SOURCE_DIR}/src)
set(OUTPUT_BINARY ${CMAKE_CURRENT_BINARY_DIR}/mecanum_control_go)

# Add custom command to build Go code
add_custom_command(
  OUTPUT ${OUTPUT_BINARY}
  COMMAND go build -o ${OUTPUT_BINARY} ${SOURCE_DIR}
  WORKING_DIRECTORY ${SOURCE_DIR}
  COMMENT "Building Go executable"
)

# Ensure the binary is built as part of the build process
add_custom_target(build_go ALL DEPENDS ${OUTPUT_BINARY})

# Install the binary to the ROS 2 standard location
install(PROGRAMS ${OUTPUT_BINARY} DESTINATION lib/${PROJECT_NAME})

# Ament package macros
ament_package()
