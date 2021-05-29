package main

import "fmt"

func LinInt(dx, y2, y1 float64) float64 {
	return 0.5*dx*(y2 + y1)
}

func Within(a, b, thresh int) bool {
	if b < (a + thresh) && b > (a - thresh) {
		return true
	}
	return false
}

func AwaitButton() {
	fmt.Scanf("%s")
}

func ButtonSig(sig *bool) {
	fmt.Scanf("%s")
	*sig = true 
}

func Modv(target int, current int, speed int) int{
	if current < (target - 10) || current > (target + 10) {
		return speed
	}
	return speed / 2
}