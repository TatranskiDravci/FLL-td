package main

import (
	"fmt"
)

func Abs(x float64) float64 {
	if x < 0 {
		return -1*x
	}
	return x
}

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
	targetAbs := Abs(target)
	currentAbs := Abs(current)

	a := -((minSpeed - maxSpeed) / (targetAbs*targetAbs))
	b := -2*a*targetAbs
	return int(a*currentAbs*currentAbs + b*currentAbs + maxSpeed)
}