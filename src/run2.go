package main


func main() {
	pid  := NewPID(0.0, 25.0, 25.0, 0.0)
	base := NewBase("A", "D", "1")
	sens1 := NewSensing("3")
    sens2 := NewSensing("2")
	shift := NewShifter("B", "C")

	// // modules
	// miso := NewM1S0(shift)
	// misi := NewM1S1(shift)
	// mise := NewM1S3(shift)

	// p1 := sens1.ColorCalib("purple_L_")
	// p2 := sens1.ColorCalib("white_L_")

	AwaitButton()
}
