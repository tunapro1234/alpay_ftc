package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"os"
	"os/signal"
	"syscall"
	"time"

	nav_msgs "nfr_odom/msgs/nav_msgs/msg"

	std_msgs "nfr_odom/msgs/std_msgs/msg"

	"github.com/tiiuae/rclgo/pkg/rclgo"
)

// Function to simulate getting encoder values from topics
func getEncoderValue(topic string) float64 {
	// Placeholder for actual encoder reading logic from ROS topics
	// Replace with actual subscription to encoder topic
	return 0.02 // Example value for demonstration
}

func run() error {
	rclArgs, _, err := rclgo.ParseArgs(os.Args[1:])
	if err != nil {
		return fmt.Errorf("failed to parse ROS args: %v", err)
	}

	if err := rclgo.Init(rclArgs); err != nil {
		return fmt.Errorf("failed to initialize rclgo: %v", err)
	}
	defer rclgo.Uninit()

	node, err := rclgo.NewNode("odometry_publisher", "")
	if err != nil {
		return fmt.Errorf("failed to create node: %v", err)
	}
	defer node.Close()

	odomPub, err := nav_msgs.NewOdometryPublisher(node, "/odom", nil)
	if err != nil {
		return fmt.Errorf("failed to create odometry publisher: %v", err)
	}
	defer odomPub.Close()

	// Load odometry configuration from the YAML file
	config, err := LoadOdometryConfig("odometry_config.yaml")
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	// if config.UseSimTime {
	// 	node.SetUseSimTime(true)
	// }

	// Example initial pose (x, y, theta)
	var x, y, theta float64

	// Set up the rates for publishing and reading encoder values
	publishInterval := time.Duration(1e9 / config.PublishRate) // Nanoseconds
	readingInterval := time.Duration(1e9 / config.ReadingRate) // Nanoseconds

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	ticker := time.NewTicker(readingInterval)
	defer ticker.Stop()

	lastPublishTime := time.Now()

	for {
		select {
		case <-ctx.Done():
			return nil

		case <-ticker.C:
			// Read encoder values from the topics defined in the config
			leftWheelDist := getEncoderValue(config.Wheels.LeftWheel.PublishTopic)
			rightWheelDist := getEncoderValue(config.Wheels.RightWheel.PublishTopic)
			perpendicularWheelDist := getEncoderValue(config.Wheels.PerpendicularWheel.PublishTopic)

			// Calculate odometry from wheel encoder distances
			dx, dy, dtheta := CalculateOdometry(config, leftWheelDist, rightWheelDist, perpendicularWheelDist)

			// Update the pose (x, y, theta) using the odometry values
			x, y, theta = UpdatePose(x, y, theta, dx, dy, dtheta)

			// Prepare odometry message to publish
			odomMsg := nav_msgs.NewOdometry()
			odomMsg.Header = *std_msgs.NewHeader()
			odomMsg.Header.FrameId = config.OdomPublishTopic
			odomMsg.ChildFrameId = config.ChildFrameId
			odomMsg.Pose.Pose.Position.X = x
			odomMsg.Pose.Pose.Position.Y = y
			odomMsg.Pose.Pose.Orientation.Z = math.Sin(theta / 2)
			odomMsg.Pose.Pose.Orientation.W = math.Cos(theta / 2)

			odomMsg.Twist.Twist.Linear.X = dx
			odomMsg.Twist.Twist.Linear.Y = dy
			odomMsg.Twist.Twist.Angular.Z = dtheta

			// Publish odometry message at the defined publish rate
			if time.Since(lastPublishTime) >= publishInterval {
				node.Logger().Infof("Publishing Odometry: x=%.2f, y=%.2f, theta=%.2f", x, y, theta)
				if err := odomPub.Publish(odomMsg); err != nil {
					node.Logger().Errorf("Failed to publish odometry: %v", err)
				}
				lastPublishTime = time.Now() // Update last publish time
			}
		}
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
