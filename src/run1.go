package main

import (
	"time"
)

func main() {
	pid  := NewPID(0.0, 25.0, 25.0, 0.0)
	base := NewBase("A", "D", "1")
	sensL := NewSensing("3")
	sensR := NewSensing("4")
	shift := NewShifter("B", "C")

	module := NewModule(shift)

	purple := sensR.ColorCalib("purpleR")
	white := sensL.ColorCalib("whiteL")

	AwaitButton()

	base.Move(400, sensR, pid, purple)
	base.RotateR(-42, 200)

	module.BeginDrive(500, 0)	// async retraction

	base.Move(100, sensL, pid, white)

	module.AwaitDrive()			// await retraction
	module.Drive(-1700, 0)		// extend
	module.Drive( 1700, 0)		// retract

	shift.BeginShifting(1)		// async shift to module 1
	base.Rotate(12, 200)
	base.MoveTimed(300, 0.39, pid)
	base.Rotate(-12, 200)

	module.Drive(900, 1)

	base.Rotate(12, 200)

	module.Drive(-1700, 0)		// extend

	shift.BeginShifting(3)		// async shift to module 3
	base.MoveTimedReverse(300, 0.61, pid)

	module.Drive(600, 3)		// open

	shift.BeginShifting(0)		// async shift to module <0, next run>
	time.Sleep(2 * time.Second)
	base.MoveTimedReverse(500, 2.0, pid)
}
