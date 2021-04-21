package main

import (
	// "github.com/ev3go/ev3"
	// "github.com/ev3go/ev3dev"
	// "github.com/ev3go/ev3dev/fb"
	// "fmt"
)

func main() {
	// initialize required modules
	// r := InitRobot("A", "D", "2", "3")
	s := InitShifter("B", "C", -150, 200)
	c := InitCarrier(s, 0, 5) // 1st mod.
	l := InitLift(s, 3) // 4th mod.

	// r.Move(1000, 350, 10, 5) // target: 1m, error: 1cm
	// r.Rotate(90, 250) // target: 90 degs.
	// r.Move(500, 350, 10, 5) // target: 0.5m, error: 1cm

	c.Release(2, 200) // drop: 2 bricks
	l.To(13.5, 200) // target: 13.5cm
	c.Release(8, 200) // drop: 8 (or remaining) bricks
	l.To(10, 200) // target: 7cm

	s.To(0) // target: 1st mod.
	for true {}
}
