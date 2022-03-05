package main

import (
	"fmt"
)

func main() {
	sens := NewSensing("4")
	col := sens.ColorCalib("purple")
	// pid  := NewPID(0.0, 25.0, 25.0, 0.0)
	// base := NewBase("A", "D", "1")
	//
	//
	// for i := 200; i <= 500; i += 50 {
	// 	fmt.Println(i)
	// 	AwaitButton()
	// 	base.Move(i, sens, pid, col)
	// }

	for {
		AwaitButton()
		fmt.Println(sens.CompareColor(col))
	}
}
