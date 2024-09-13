package main

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Wheel struct {
	Diameter     float64    `yaml:"diameter"`
	Position     [2]float64 `yaml:"position"`
	PublishTopic string     `yaml:"publish_topic"`
	EncoderCPR   int        `yaml:"encoder_cpr"`
}

type OdometryConfig struct {
	UseSimTime       bool   `yaml:"use_sim_time"`
	ChildFrameId     string `yaml:"child_frame_id"`
	OdomPublishTopic string `yaml:"odom_publish_topic"`
	PublishRate      int    `yaml:"publish_rate"`
	ReadingRate      int    `yaml:"reading_rate"`
	Wheels           struct {
		LeftWheel          Wheel `yaml:"left_wheel"`
		RightWheel         Wheel `yaml:"right_wheel"`
		PerpendicularWheel Wheel `yaml:"perpendicular_wheel"`
	} `yaml:"wheels"`
}

// LoadOdometryConfig reads the YAML configuration file and unmarshals it into the OdometryConfig struct
func LoadOdometryConfig(filename string) (*OdometryConfig, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read YAML file: %v", err)
	}

	var config OdometryConfig
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to parse YAML file: %v", err)
	}

	if config.Wheels.LeftWheel.Position[0] < config.Wheels.RightWheel.Position[1] {
		return nil, fmt.Errorf("left wheel should be on the left of the right wheel")
	}
	return &config, nil
}
