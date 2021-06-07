package main

func main() {
	r := InitRobot("A", "D", "2", "4", "1", "3")
	R := r
	R.colorLeft = r.colorMid
	ClearKeys()
	r.ColorCalib("GREEN")
	r.ColorCalib("PINK")
	r.ColorCalib("BLACK")
	r.ColorCalib("WHITE")
	r.ColorCalib("GREY")
	r.ColorCalib("YELLOW")
	R.ColorCalib("MWHITE")
	R.ColorCalib("MBLACK")
	R.ColorCalib("MGREY")
}