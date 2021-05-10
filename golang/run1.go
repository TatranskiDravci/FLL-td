package main

func main() {
	// initialize required modules
	r := InitRobot("A", "D", "2", "4")
	// path code
	r.Run(1800, 900)
	r.Run(-1900, 700)
}