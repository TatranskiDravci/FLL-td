package main

import (
	"time"
)

/*
	basic PID interface
		SP	- setpoint
		K_P - proportional	parameter
		K_I - integral		parameter
		K_D - differential	parameter
*/
type PID struct {
	SP  float64
	K_P	float64
	K_I float64
	K_D	float64
}

/*
	basic PID constructor
		SP	- setpoint
		K_P - proportional	parameter
		K_I - integral		parameter
		K_D - differential	parameter
*/
func NewPID(SP, K_P, K_I, K_D float64) PID {
	return PID {
		SP:   SP,
		K_P: K_P,
		K_I: K_I,
		K_D: K_D,
	}
}

/*
	non-stop PID interface
		SP	- setpoint
		K_P - proportional	parameter
		K_I - integral		parameter
		K_D - differential	parameter

		integral	- transfered integral value
		perror		- transfered preceeding error value
		ptime		- transfered preceeding time
		start		- transfered time on Move initialization

*/
type PIDNS struct {
	// storage block
	SP  float64
	K_P	float64
	K_I float64
	K_D	float64
	// transfer block
	integral	float64
	perror 		float64
	ptime  		float64
	start		time.Time
}

/*
	non-stop PID constructor (PID -> PIDNS copy constructor)
		ctl - basic PID object
*/
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
