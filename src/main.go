package main


func main() {
	// black := [3]int{26, 16, 40}
 	pink := [3]int{184, 94, 110}

	r := InitRobot("A", "D", "2", "4", "1", "3")


	// r.ColorPrint()
	r.Move(400, pink, 20, 28, 40, 1)
}