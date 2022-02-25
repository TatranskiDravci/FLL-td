package main

func main() {
	sens1 := NewSensing("3")
	// sens2 := NewSensing("2")
	// ClearAll()
	sens1.ColorCalib("black")
	sens1.ColorCalib("green")
}
