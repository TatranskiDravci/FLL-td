package main


func main() {
	pid  := NewPID(0.0, 25.0, 25.0, 0.0)
    pidns := NewPIDNS(pid)
	base := NewBase("A", "D", "1")

	// sensL := NewSensing("3")
    sensB := NewSensing("2")
	// sensR := NewSensing("4")

	shift := NewShifter("B", "C")
	module := NewModule(shift)

	// colors
	drabB := sensB.ColorCalib("drabB")
	blackB := sensB.ColorCalib("blackB")
	whiteB := sensB.ColorCalib("whiteB")
	blueB := sensB.ColorCalib("blueB")
	//
	// greyL := sensL.ColorCalib("greyL")
	// greenL := sensL.ColorCalib("greenL")
	// whiteL := sensL.ColorCalib("whiteL")
	//
	// blackR := sensR.ColorCalib("blackR")


	// run
	AwaitButton()

	base.MoveNS(300, sensB, &pidns, drabB, _initial)
	base.MoveNS(300, sensB, &pidns, whiteB, _final)

	base.RotateR(-40, 200)
	module.AwaitDrive()

	base.MoveTimed(300, 0.7, pid)


	base.Move(300, sensB, pid, whiteB)
	base.Move(300, sensB, pid, blackB)
	module.Drive(-4*360, 0)
	base.MoveNS(300, sensB, &pidns, blueB, _initial)
	base.MoveNS(300, sensB, &pidns, blackB, _final)
}
