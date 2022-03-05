package main

import (
	"time"
)

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

type PIDNS struct {
	// storage block
	K_P	float64
	K_I float64
	K_D	float64
	SP  float64
	// functional block
	integral	float64
	perror 		float64
	ptime  		float64
	start		time.Time
}

func NewPIDNS(ctl PID) PIDNS {
	return PIDNS {
		SP:   ctl.SP,
		K_P: ctl.K_P,
		K_I: ctl.K_I,
		K_D: ctl.K_D,
		perror:		0.0,
		ptime: 		0.0,
		integral: 	0.0,
		start:		time.Time{},
	}
}
