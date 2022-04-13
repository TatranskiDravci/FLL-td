package main


func main() {
	pid   := NewPID(0.0, 25.0, 25.0, 0.0)
	pid15 := NewPID(15.0, 5.0, 5.0, 0.25)
    // pidns := NewPIDNS(pid)
	i := NewPID(0.0, 1.0, 10.0, 0.0)
	// ins := NewPIDNS(i)
	base := NewBase("A", "D", "1")

	// sensL := NewSensing("3")
    sensB := NewSensing("2")
	// sensR := NewSensing("4")

	shift := NewShifter("B", "C")
	module := NewModule(shift)

	// colors
	// drabB := sensB.ColorCalib("drabB")
	// blackB := sensB.ColorCalib("blackB")
	whiteB := sensB.ColorCalib("whiteB")
	// blueB  := sensB.ColorCalib("blueB")
	//
	// greyL := sensL.ColorCalib("greyL")
	// greenL := sensL.ColorCalib("greenL")
	// whiteL := sensL.ColorCalib("whiteL")
	//
	// blackR := sensR.ColorCalib("blackR")


	// run
	AwaitButton()

	// base.MoveTimedNS(300, 0.5, &pidns, _initial)
	// base.MoveNS(300, sensB, &pidns, drabB, _initial)
	// base.MoveNS(300, sensB, &pidns, whiteB, _final)
	//
	// base.RotateR(-45, 300)
	// module.AwaitDrive()
	//
	//
	//
	// base.MoveNS(300, sensB, &pidns, whiteB, _initial)
	// base.MoveNS(300, sensB, &pidns, blackB, _final)
	// module.Drive(-4*360, 0)
	// base.MoveNS(300, sensB, &pidns, blueB, _initial)
	// base.MoveNS(300, sensR, &pidns, blackR, _final)
	// module.Drive(-600, 0)
	// module.Drive(4*360 + 600, 0)
	//
	shift.BeginShifting(2)
	base.Rotate(-15, 100)
	base.MoveTimed(300, 2.0, pid15)

	// base.MoveTimed(300, 0.1, pid)
	module.Drive(-4*360, 2)

	shift.BeginShifting(3)
	base.Rotate(-10, 100)

	// base.MoveReverseNS(300, sensB, &pidns, blackB, _initial)
	// base.MoveReverseNS(300, sensB, &pidns, whiteB, _midrun)
	base.MoveTimedReverse(300, 0.8, pid)

	base.RotateL(100, 300)
	module.Drive(1.5*360, 3)
	base.MoveReverse(300, sensB, i, whiteB)
	base.MoveTimed(300, 0.1, i)
	module.Drive(-1.5*360, 3)
}
