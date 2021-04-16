package main

import (
	// "github.com/ev3go/ev3"
	// "github.com/ev3go/ev3dev"
	// "github.com/ev3go/ev3dev/fb"
	// "fmt"
)

func main() {
	r := InitRobot("A", "D", "2", "3")
	s := InitShifter("B", "C", -150, 200)
	l := InitLift(s, 3)
	r.Move(350, 1000, 10, 5)
	r.Rotate(250, 90)
	r.Move(350, 500, 10, 5)
	l.To(13.5, 200)
	l.To(7, 200)
	l.To(9, 200)
	l.To(7, 200)
	s.To(0)
	for true {}
}
