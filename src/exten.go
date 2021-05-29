package main

import "fmt"

func flinInt(dx, y2, y1 float64) float64 {
	return 0.5*dx*(y2 + y1)
}

func within(a, b, thresh int) bool {
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
