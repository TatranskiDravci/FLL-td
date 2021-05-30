package main

import (
	"fmt"
	"math"
)

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

func ModSpeed(target, current, minSpeed, maxSpeed float64) int{
	targetAbs := math.Abs(target)
	currentAbs := math.Abs(current)

	base := math.Pow(5*(maxSpeed - minSpeed), 1 / targetAbs)
	shift := math.Log(maxSpeed - minSpeed) / math.Log(base) 

	return int(math.Pow(base, shift - target) + minSpeed)
}