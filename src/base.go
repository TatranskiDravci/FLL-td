package main

import (
	"github.com/ev3go/ev3dev"
	"math"
	"time"
)

/*
	lm 	 - left  motor
	rm   - right motor
	gyro - gyro sensor
*/
type Base struct {
	lm   *ev3dev.TachoMotor
	rm   *ev3dev.TachoMotor
	gyro *ev3dev.Sensor
}

/*
	lp    - left  motor port
	rp 	  - right motor port
	gyrop - gyro sensor port
*/
func NewBase(lp, rp, gyrop string) Base {
	lm, _ := ev3dev.TachoMotorFor("ev3-ports:out" + lp, "lego-ev3-l-motor")
	rm, _ := ev3dev.TachoMotorFor("ev3-ports:out" + rp, "lego-ev3-l-motor")
	gyro, _ := ev3dev.SensorFor("ev3-ports:in" + gyrop, "lego-ev3-gyro")

	// reset gyro sensor on init
	gyro.SetMode("GYRO-RATE")
	gyro.SetMode("GYRO-ANG")

	return Base {
		lm: 	lm,
		rm: 	rm,
		gyro: gyro,
	}
}

// reset on-board gyroscope
func (this Base) ResetGyro() {
	this.gyro.SetMode("GYRO-RATE")
	time.Sleep(500 * time.Millisecond)
	this.gyro.SetMode("GYRO-ANG")
}

// set stop action
func (this Base) SetBrake() {
	this.lm.SetStopAction("hold")
	this.rm.SetStopAction("hold")
}

// run tank: set individual motor speeds manually
func (this Base) RunTank(ls, rs int) {
	this.lm.SetSpeedSetpoint(ls).Command("run-forever")
	this.rm.SetSpeedSetpoint(rs).Command("run-forever")
}

// run steering: set motor speeds automatically
func (this Base) RunSteering(max, x float64) {
	ls := math.Max(math.Min( x + max, max), -max)
	rs := math.Max(math.Min(-x + max, max), -max)
	this.lm.SetSpeedSetpoint(int(ls)).Command("run-forever")
	this.rm.SetSpeedSetpoint(int(rs)).Command("run-forever")
}

// run steering: set motor speeds automatically
func (this Base) RunSteeringReverse(max, x float64) {
	ls := math.Max(math.Min( x + max, max), -max)
	rs := math.Max(math.Min(-x + max, max), -max)
	this.lm.SetSpeedSetpoint(-int(ls)).Command("run-forever")
	this.rm.SetSpeedSetpoint(-int(rs)).Command("run-forever")
}

// stop motors
func (this Base) Stop() {
	this.rm.Command("stop")
	this.lm.Command("stop")
}
