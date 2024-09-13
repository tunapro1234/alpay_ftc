// src/config_parser.cpp

#include "nfr_odom/config_parser.hpp"

OdometryConfigParser::OdometryConfigParser() = default;

void OdometryConfigParser::loadConfig(const std::string &config_path)
{
  try
  {
    YAML::Node config = YAML::LoadFile(config_path);

    // Parse basic parameters
    use_sim_time_ = config["use_sim_time"].as<bool>();
    publish_rate_ = config["publish_rate"].as<int>();
    reading_rate_ = config["reading_rate"].as<int>();
    child_frame_id_ = config["child_frame_id"].as<std::string>();
    odom_publish_topic_ = config["odom_publish_topic"].as<std::string>();

    // Parse wheels configuration
    const YAML::Node &wheels = config["wheels"];
    parseWheelConfig(wheels["left_wheel"], "left_wheel");
    parseWheelConfig(wheels["right_wheel"], "right_wheel");
    parseWheelConfig(wheels["perpendicular_wheel"], "perpendicular_wheel");

    RCLCPP_INFO(rclcpp::get_logger("OdometryConfigParser"), "Configuration loaded successfully.");
  }
  catch (const YAML::Exception &e)
  {
    RCLCPP_ERROR(rclcpp::get_logger("OdometryConfigParser"), "Failed to load configuration: %s", e.what());
  }
}

void OdometryConfigParser::parseWheelConfig(const YAML::Node &wheel_node, const std::string &wheel_name)
{
  WheelConfig wheel;
  wheel.diameter = wheel_node["diameter"].as<double>();
  wheel.position = wheel_node["position"].as<std::vector<double>>();
  wheel.publish_topic = wheel_node["publish_topic"].as<std::string>();
  wheel.encoder_cpr = wheel_node["encoder_cpr"].as<int>();

  wheels_[wheel_name] = wheel;
}

// Getters
bool OdometryConfigParser::getUseSimTime() const { return use_sim_time_; }
int OdometryConfigParser::getPublishRate() const { return publish_rate_; }
int OdometryConfigParser::getReadingRate() const { return reading_rate_; }
std::string OdometryConfigParser::getChildFrameId() const { return child_frame_id_; }
std::string OdometryConfigParser::getOdomPublishTopic() const { return odom_publish_topic_; }
const std::map<std::string, WheelConfig>& OdometryConfigParser::getWheels() const { return wheels_; }
