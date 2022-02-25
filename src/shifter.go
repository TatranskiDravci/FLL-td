package main

import "github.com/ev3go/ev3dev"

const (
	_optimalRate = 400
	_offset = 90
)

type Shifter struct {
	sm *ev3dev.TachoMotor
	dm *ev3dev.TachoMotor
	nullstate ev3dev.MotorState
}


func NewShifter(sp, dp string) Shifter {
	sm, _ := ev3dev.TachoMotorFor("ev3-ports:out" + sp, "lego-ev3-l-motor")
	dm, _ := ev3dev.TachoMotorFor("ev3-ports:out" + dp, "lego-ev3-l-motor")
	sm.SetPosition(0)
	dm.SetPosition(0)
	nullstate, _ := sm.State()
	return Shifter {
		nullstate: nullstate,
		sm:	sm,
		dm: dm,
	}
}

/*
	asynchronous shifting
		target	- output number
*/
func (this Shifter) BeginShifting(target int) {
	this.nullstate, _ = this.sm.State()
	this.sm.SetStopAction("brake")
	this.sm.SetSpeedSetpoint(_optimalRate)
	this.sm.SetPositionSetpoint(_offset * target).Command("run-to-abs-pos")
}

/*
	await clause for shifting, should be called (at some point) after any BaginShifting call
*/
func (this Shifter) AwaitShifting() {
	for state, _ := this.sm.State(); state != this.nullstate; state, _ = this.sm.State() {}
}

/*
	upper level function for relative module control (synchronous)
		target	- relative target angle
		rate	- turn rate of the motor (I think??)
*/
func (this Shifter) DriveRelative(target, rate int) {
	this.nullstate, _ = this.dm.State()
	this.dm.SetStopAction("brake")
	this.dm.SetSpeedSetpoint(rate)
	this.dm.SetPositionSetpoint(target).Command("run-to-rel-pos")
	for state, _ := this.dm.State(); state != this.nullstate; state, _ = this.dm.State() {}
}

/*
	upper level function for absolute module control (synchronous)
		target	- absolute target angle
		rate	- turn rate of the motor (I think??)
*/
func (this Shifter) DriveAbsolute(target, rate int) {
	this.nullstate, _ = this.dm.State()
	this.dm.SetStopAction("brake")
	this.dm.SetSpeedSetpoint(rate)
	this.dm.SetPositionSetpoint(target).Command("run-to-abs-pos")
	for state, _ := this.dm.State(); state != this.nullstate; state, _ = this.dm.State() {}
}

func (this Shifter) BeginDriveRelative(target, rate int) {
	this.nullstate, _ = this.dm.State()
	this.dm.SetStopAction("brake")
	this.dm.SetSpeedSetpoint(rate)
	this.dm.SetPositionSetpoint(target).Command("run-to-rel-pos")
}

func (this Shifter) AwaitDriveRelative() {
	for state, _ := this.dm.State(); state != this.nullstate; state, _ = this.dm.State() {}
}
