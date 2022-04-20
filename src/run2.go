package main


func main() {
	pid   := NewPID(0.0, 25.0, 25.0, 0.0)
	null  := NewPID(0.0, 0, 0, 0.0)
	// pid15 := NewPID(15.0, 5.0, 5.0, 0.25)
    pidns := NewPIDNS(pid)
	// i := NewPID(0.0, 1.0, 10.0, 0.0)
	// ins := NewPIDNS(i)
	base := NewBase("A", "D", "1")

	// sensL := NewSensing("3")
    sensB := NewSensing("2")
	sensR := NewSensing("4")

	shift := NewShifter("B", "C")
	module := NewModule(shift)

	// colors
	// drabB := sensB.ColorCalib("drabB")
	// blackB := sensB.ColorCalib("blackB")
	// whiteB := sensB.ColorCalib("whiteB")
	blueB  := sensB.ColorCalib("blueB")
	//
	// greyL := sensL.ColorCalib("greyL")
	// greenL := sensL.ColorCalib("greenL")
	// whiteL := sensL.ColorCalib("whiteL")
	//
	// blackR := sensR.ColorCalib("blackR")
	whiteR := sensR.ColorCalib("whiteR")


	// run
	AwaitButton()

	// SECTOR 1
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

	// SECTOR 2
	shift.BeginShifting(2)
	base.MoveTimedNS(300, 0.6, &pidns, _initial)
	base.MoveNS(300, sensR, &pidns, whiteR, _final)

	// base.MoveTimed(300, 0.1, pid)
	module.Drive(-4*360, 2)

	shift.BeginShifting(3)
	// base.Rotate(-10, 100)

	// base.MoveReverseNS(300, sensB, &pidns, blackB, _initial)
	// base.MoveReverseNS(300, sensB, &pidns, whiteB, _midrun)
	base.MoveTimedReverse(200, 1.35, pid)

	base.RotateL(90, 300)
	module.Drive(1.5*360, 3)
	base.MoveTimedReverse(300, 2.0, null)
	base.MoveTimed(300, 0.1, pid)
	module.Drive(-2.25*360, 3)

	// SECTOR 3
	shift.BeginShifting(1)
	base.RotateR( 90, 200)
	module.Drive(-8*360, 1)
	module.Drive( 8*360, 1)
	shift.BeginShifting(0)
	base.RotateR(-90, 200)
	base.Move(300, sensR, pid, blueB)
}
