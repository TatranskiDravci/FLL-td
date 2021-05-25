package main

import (
	"fmt"
)

func main() {
	// r := InitRobot("A", "D", "2", "4", "1", "3")
	s := InitShifter("B", "C", -1800, 550)
	// pink := r.ColorCalib("pink")
	// black := r.ColorCalib("black")

	for i := 0; i < 5; i++ {
		s.To(i)
		s.Run(520, 400)
		var button string
		fmt.Scanf("%s", &button)
	}
	s.To(0)
}