package main

func main() {
	// sens1 := NewSensing("3")
	// // sens2 := NewSensing("2")
	// ClearAll()
	// sens1.ColorCalib("whiteL")
	sensL := NewSensing("3")
    sensB := NewSensing("2")
	sensR := NewSensing("4")

	sensB.ColorCalib("drabB")
	sensB.ColorCalib("blackB")
	sensB.ColorCalib("whiteB")
	sensB.ColorCalib("blueB")

	sensL.ColorCalib("greyL")
	sensL.ColorCalib("greenL")
	sensL.ColorCalib("whiteL")

	sensR.ColorCalib("blackR")

}
