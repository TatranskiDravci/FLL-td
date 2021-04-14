package main

import (
	// "github.com/ev3go/ev3"
	// "github.com/ev3go/ev3dev"
	// "github.com/ev3go/ev3dev/fb"
)

func main() {
	r := InitRobot("A", "D", "2", "3")
	r.Move(350, 1000, 10, 5)
	r.Rotate(250, 90)
}