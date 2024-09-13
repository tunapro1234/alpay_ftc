// include/nfr_odom/config_parser.hpp

#ifndef __NFR_ODOM_CONFIG_PARSER_HPP__
#define __NFR_ODOM_CONFIG_PARSER_HPP__

#include <rclcpp/rclcpp.hpp>
#include <yaml-cpp/yaml.h>
#include <string>
#include <vector>
#include <map>

struct WheelConfig {
  double diameter;
  std::vector<double> position;
  std::string publish_topic;
  int encoder_cpr;
};

class OdometryConfigParser {
public:
  OdometryConfigParser();
  void loadConfig(const std::string &config_path);

  // Getters
  bool getUseSimTime() const;
  int getPublishRate() const;
  int getReadingRate() const;
  std::string getChildFrameId() const;
  std::string getOdomPublishTopic() const;
  const std::map<std::string, WheelConfig>& getWheels() const;

private:
  void parseWheelConfig(const YAML::Node &wheel_node, const std::string &wheel_name);

  // Configuration variables
  bool use_sim_time_;
  int publish_rate_;
  int reading_rate_;
  std::string child_frame_id_;
  std::string odom_publish_topic_;
  std::map<std::string, WheelConfig> wheels_;
};

#endif // __NFR_ODOM_CONFIG_PARSER_HPP__
