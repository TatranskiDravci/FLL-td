package main

import "os"

func main() {
	r := InitRobot("A", "D", "2", "4", "1", "3")
	os.Clearenv()
	r.ColorCalib("PINK")
	r.ColorCalib("BLACK")
	r.ColorCalib("WHITE")
	r.ColorCalib("GREEN")
	r.ColorCalib("BLUE")
}