package main

import (
	"time"
)

func main() {
	pid  := NewPID(0.0, 25.0, 25.0, 0.0)
	base := NewBase("A", "D", "1")
	sens1 := NewSensing("3")
	shift := NewShifter("B", "C")

	// modules
	miso := NewM1S0(shift)
	misi := NewM1S1(shift)
	mise := NewM1S3(shift)

	sens1.ProfileCalib("1")

	p1 := sens1.ColorCalib("purple_L_")
	p2 := sens1.ColorCalib("white_L_")

	AwaitButton()

	base.Move(450, p1, sens1, pid)
	base.RotateR(-40, 200)

	miso.BeginRetract()

	base.Move(200, p2, sens1, pid)

	miso.AwaitRetract()
	miso.Extend()
	miso.Retract()

	shift.BeginShifting(1) // shift to M1S1

	base.Rotate(12, 200)
	base.MoveTimed(300, 0.360, pid)

	misi.Flip()
	shift.BeginShifting(0); shift.AwaitShifting()

	miso.Extend()

	shift.BeginShifting(3)

	base.MoveTimedReverse(300, 0.61, pid)

	mise.Open()

	time.Sleep(2 * time.Second)

	base.MoveTimedReverse(500, 2.0, pid)
}
