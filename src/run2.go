package main

import "fmt"

/*
	todo:
		- calibrate timed functions
		- run tests on sector 3 & 4 --> calibrate values
		- run full-run tests		--> calibrate landings
*/

func main() {
	// robot config
	pid   := NewPID(0.0, 25.0, 25.0, 0.0)
	null  := NewPID(0.0, 0, 0, 0.0)
	pidns := NewPIDNS(pid)
	base := NewBase("A", "D", "1")

	sensB := NewSensing("2")
	sensR := NewSensing("4")
	// sensL := NewSensing("3")

	shift := NewShifter("B", "C")
	module := NewModule(shift)

	
	// colors - uncomment required colors
	drabB := sensB.ColorCalib("drabB")
	blackB := sensB.ColorCalib("blackB")
	whiteB := sensB.ColorCalib("whiteB")
	blueB  := sensB.ColorCalib("blueB")

	blackR := sensR.ColorCalib("blackR")
	whiteR := sensR.ColorCalib("whiteR")

	// greyL := sensL.ColorCalib("greyL")
	// greenL := sensL.ColorCalib("greenL")
	// whiteL := sensL.ColorCalib("whiteL")


	// run
	fmt.Println("READY")
	AwaitButton()

	// sector 1 - uncomment if needed
	fmt.Println("SECTOR 1")
	base.MoveTimedNS(300, 0.5, &pidns, _initial)
	base.MoveNS(300, sensB, &pidns, drabB, _initial)
	base.MoveNS(300, sensB, &pidns, whiteB, _final)
	base.RotateR(-45, 300)
	base.MoveNS(300, sensB, &pidns, whiteB, _initial)
	base.MoveNS(300, sensB, &pidns, blackB, _final)
	module.Drive(-4*360, 0)								// extend module 1 @[0] for crane 
	base.MoveNS(300, sensB, &pidns, blueB, _initial)
	base.MoveNS(300, sensR, &pidns, blackR, _final)
	module.Drive(-600, 0)								// drop package (module 1 @[0])
	module.Drive(4*360 + 600, 0)						// retract module 1 @[0]
	

	// sector 2
	fmt.Println("SECTOR 2")
	shift.BeginShifting(2)								// async shift to module 3 @[2]
	base.MoveTimedNS(300, 0.6, &pidns, _initial)
	base.MoveNS(300, sensR, &pidns, whiteR, _final)
	module.Drive(-4*360, 2)								// drop cargo (module 3 @[2])
	shift.BeginShifting(3)								// async shift to module 4 @[3] 
	base.MoveTimedReverse(200, 1.35, pid)
	base.RotateL(90, 300)
	module.Drive(1.5*360, 3)							// lower module 4 @[3]
	base.MoveTimedReverse(300, 2.0, null)
	base.MoveTimed(300, 0.1, pid)
	module.Drive(-2.25*360, 3)							// raise module 4 @[3]

	// sector 3
	fmt.Println("SECTOR 3")
	shift.BeginShifting(1)								// async shift to module 2 @[1]
	base.RotateR( 90, 200)
	module.Drive(-8*360, 1)								// lower module 2 @[1]
	module.Drive( 8*360, 1)								// raise module 2 @[1]
	shift.BeginShifting(0)								// async shift to module 1 @[0]
	base.RotateR(-90, 200)
	base.Move(300, sensR, pid, blueB)
	module.Drive(-4*360, 0)
	base.MoveTimed(300, 1.0, pid)						// todo: calibrate
	base.MoveTimedReverse(300, 1.0, pid)				// todo: calibrate
	
	// sector 4
	fmt.Println("SECTOR 4")
	base.Rotate(90, 300)
	base.MoveTimedReverse(500, 1.0, pid)				// todo: calibrate
}
