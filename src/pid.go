package main

type PID struct {
	K_P	float64
	K_I float64
	K_D	float64
	SP  float64
}

func NewPID(SP, K_P, K_I, K_D float64) PID {
	return PID {
		SP:   SP,
		K_P: K_P,
		K_I: K_I,
		K_D: K_D,
	}
}
