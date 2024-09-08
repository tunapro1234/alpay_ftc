package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	geometry_msgs_msg "mecanum_control_go/msgs/geometry_msgs/msg"
	std_msgs_msg "mecanum_control_go/msgs/std_msgs/msg"

	"github.com/tiiuae/rclgo/pkg/rclgo"
)

// WheelVelocities holds velocities for the four mecanum wheels.
type WheelVelocities struct {
	FrontLeft  float64
	FrontRight float64
	RearLeft   float64
	RearRight  float64
}

// ConvertTwistToWheelVelocities converts a Twist message to individual wheel velocities.
func ConvertTwistToWheelVelocities(twist *geometry_msgs_msg.Twist) WheelVelocities {
	// Robot parameters
	wheelRadius := 0.05
	l := 0.5 // length between wheels
	w := 0.3 // width between wheels

	// Linear and angular velocities from the Twist message
	vx := twist.Linear.X
	vy := twist.Linear.Y
	vtheta := twist.Angular.Z

	// Apply mecanum kinematics
	fl := (vx - vy - vtheta*(l+w)) / wheelRadius
	fr := (vx + vy + vtheta*(l+w)) / wheelRadius
	rl := (vx + vy - vtheta*(l+w)) / wheelRadius
	rr := (vx - vy + vtheta*(l+w)) / wheelRadius

	return WheelVelocities{FrontLeft: fl, FrontRight: fr, RearLeft: rl, RearRight: rr}
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

	node, err := rclgo.NewNode("mecanum_controller", "")
	if err != nil {
		return fmt.Errorf("failed to create node: %v", err)
	}
	defer node.Close()

	// Publishers for each wheel velocity
	pubFL, err := std_msgs_msg.NewFloat64Publisher(node, "front_left_wheel_velocity", nil)
	if err != nil {
		return fmt.Errorf("failed to create front left wheel publisher: %v", err)
	}
	defer pubFL.Close()

	pubFR, err := std_msgs_msg.NewFloat64Publisher(node, "front_right_wheel_velocity", nil)
	if err != nil {
		return fmt.Errorf("failed to create front right wheel publisher: %v", err)
	}
	defer pubFR.Close()

	pubRL, err := std_msgs_msg.NewFloat64Publisher(node, "rear_left_wheel_velocity", nil)
	if err != nil {
		return fmt.Errorf("failed to create rear left wheel publisher: %v", err)
	}
	defer pubRL.Close()

	pubRR, err := std_msgs_msg.NewFloat64Publisher(node, "rear_right_wheel_velocity", nil)
	if err != nil {
		return fmt.Errorf("failed to create rear right wheel publisher: %v", err)
	}
	defer pubRR.Close()

	// Subscriber for the /cmd_vel topic
	sub, err := geometry_msgs_msg.NewTwistSubscription(
		node,
		"/cmd_vel",
		nil,
		func(msg *geometry_msgs_msg.Twist, info *rclgo.MessageInfo, err error) {
			if err != nil {
				node.Logger().Errorf("failed to receive message: %v", err)
				return
			}
			velocities := ConvertTwistToWheelVelocities(msg)

			// Create Float64 messages to publish
			flMsg := std_msgs_msg.NewFloat64()
			frMsg := std_msgs_msg.NewFloat64()
			rlMsg := std_msgs_msg.NewFloat64()
			rrMsg := std_msgs_msg.NewFloat64()

			flMsg.Data = velocities.FrontLeft
			frMsg.Data = velocities.FrontRight
			rlMsg.Data = velocities.RearLeft
			rrMsg.Data = velocities.RearRight

			// Publish to each wheel topic
			if err := pubFL.Publish(flMsg); err != nil {
				node.Logger().Errorf("Failed to publish front left wheel velocity: %v", err)
			}
			if err := pubFR.Publish(frMsg); err != nil {
				node.Logger().Errorf("Failed to publish front right wheel velocity: %v", err)
			}
			if err := pubRL.Publish(rlMsg); err != nil {
				node.Logger().Errorf("Failed to publish rear left wheel velocity: %v", err)
			}
			if err := pubRR.Publish(rrMsg); err != nil {
				node.Logger().Errorf("Failed to publish rear right wheel velocity: %v", err)
			}
		})

	if err != nil {
		return fmt.Errorf("failed to create subscriber: %v", err)
	}
	defer sub.Close()

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	// Run the node until it's interrupted
	for {
		select {
		case <-ctx.Done():
			node.Logger().Info("Shutting down mecanum controller")
			return nil
		case <-time.After(time.Second):
			// You can add periodic logging or checks here
		}
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
