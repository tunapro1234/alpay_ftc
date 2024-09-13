package main

import (
	"math"
)

// Function to calculate odometry (dx, dy, dtheta)
func CalculateOdometry(config *OdometryConfig, leftWheelDist, rightWheelDist, perpendicularWheelDist float64) (float64, float64, float64) {
	// Wheel positions (X, Y relative to the center of the robot)
	leftPos := config.Wheels.LeftWheel.Position
	rightPos := config.Wheels.RightWheel.Position

	// Distance between wheels and center
	wheelBase := rightPos[1] - leftPos[1] // Distance between left and right wheels

	// Calculate changes in position and orientation
	dx := (leftWheelDist + rightWheelDist) / 2             // Forward/backward movement
	dy := perpendicularWheelDist                           // Sideways movement
	dtheta := (rightWheelDist - leftWheelDist) / wheelBase // Rotation around the Z axis (theta)

	return dx, dy, dtheta
}

// Function to calculate the pose (x, y, theta) from odometry
func UpdatePose(x, y, theta, dx, dy, dtheta float64) (float64, float64, float64) {
	// Update position based on odometry calculations
	newX := x + dx*math.Cos(theta) - dy*math.Sin(theta)
	newY := y + dx*math.Sin(theta) + dy*math.Cos(theta)
	newTheta := theta + dtheta

	return newX, newY, newTheta
}
