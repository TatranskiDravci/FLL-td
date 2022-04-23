package main

import "github.com/ev3go/ev3dev"

const (
	_optimalRate = 400			// optimal shifter rate
	_offset = 90				// offset on shifter
)

/*
    Shifter interface
        sm			- shifter motor
		dm 			- driver  motor
		nullstate	- state id of a stopped motor
*/
type Shifter struct {
	sm *ev3dev.TachoMotor
	dm *ev3dev.TachoMotor
	nullstate ev3dev.MotorState
}

/*
    Shifter constructor
        sp	- shifter motor port
		dp 	- driver  motor port
*/
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
	provides asynchronous Shifting
		target	- output number
*/
func (this Shifter) BeginShifting(target int) {
	this.nullstate, _ = this.sm.State()
	this.sm.SetStopAction("brake")
	this.sm.SetSpeedSetpoint(_optimalRate)
	this.sm.SetPositionSetpoint(_offset * target).Command("run-to-abs-pos")
}

/*
    provides await "operator" for asynchronous Shifting
*/
func (this Shifter) AwaitShifting() {
	for state, _ := this.sm.State(); state != this.nullstate; state, _ = this.sm.State() {}
}

/*
	provides function for synchronous relative module control
		target	- relative target angle
		rate	- turn rate of the motor
*/
func (this Shifter) DriveRelative(target, rate int) {
	this.nullstate, _ = this.dm.State()
	this.dm.SetStopAction("brake")
	this.dm.SetSpeedSetpoint(rate)
	this.dm.SetPositionSetpoint(target).Command("run-to-rel-pos")
	for state, _ := this.dm.State(); state != this.nullstate; state, _ = this.dm.State() {}
}

/*
	provides function for synchronous absolute module control
		target	- absolute target angle
		rate	- turn rate of the motor
*/
func (this Shifter) DriveAbsolute(target, rate int) {
	this.nullstate, _ = this.dm.State()
	this.dm.SetStopAction("brake")
	this.dm.SetSpeedSetpoint(rate)
	this.dm.SetPositionSetpoint(target).Command("run-to-abs-pos")
	for state, _ := this.dm.State(); state != this.nullstate; state, _ = this.dm.State() {}
}

/*
	provides function for asynchronous relative module control
		target	- relative target angle
		rate	- turn rate of the motor
*/
func (this Shifter) BeginDriveRelative(target, rate int) {
	this.nullstate, _ = this.dm.State()
	this.dm.SetStopAction("brake")
	this.dm.SetSpeedSetpoint(rate)
	this.dm.SetPositionSetpoint(target).Command("run-to-rel-pos")
}

/*
    provides await "operator" for asynchronous relative module control
*/
func (this Shifter) AwaitDriveRelative() {
	for state, _ := this.dm.State(); state != this.nullstate; state, _ = this.dm.State() {}
}
