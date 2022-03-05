package main

import (
	"time"
)

func main() {
	pid  := NewPID(0.0, 25.0, 25.0, 0.0)
	base := NewBase("A", "D", "1")
	sens1 := NewSensing("3")
	sens2 := NewSensing("4")
	shift := NewShifter("B", "C")

	// modules
	miso := NewM1S0(shift)
	misi := NewM1S1(shift)
	mise := NewM1S3(shift)

	p1 := sens2.ColorCalib("purple")
	p2 := sens1.ColorCalib("white")

	AwaitButton()

	base.Move(400, sens2, pid, p1)
	base.RotateR(-42, 200)

	miso.BeginRetract()

	base.Move(100, sens1, pid, p2)

	miso.AwaitRetract()
	miso.Extend()
	miso.Retract()

	shift.BeginShifting(1) // shift to M1S1

	base.Rotate(12, 200)
	base.MoveTimed(300, 0.39, pid)
	base.Rotate(-12, 200)

	misi.Flip()
	base.Rotate(12, 200)
	shift.BeginShifting(0); shift.AwaitShifting()

	miso.Extend()

	shift.BeginShifting(3)

	base.MoveTimedReverse(300, 0.61, pid)

	mise.Open()

	time.Sleep(2 * time.Second)

	base.MoveTimedReverse(500, 2.0, pid)
}
