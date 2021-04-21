package main

import (
	// "github.com/ev3go/ev3"
	"github.com/ev3go/ev3dev"
	// "github.com/ev3go/ev3dev/fb"
	// "math"
	// "time"
	// "fmt"
)

type Shifter struct {
	shifterMotor	*ev3dev.TachoMotor
	runnerMotor		*ev3dev.TachoMotor
	offset			int
	rate			int
}

func InitShifter(shifterPort string, runnerPort string, offsetVal int, shiftrt int) Shifter {
	shifter,	_ := ev3dev.TachoMotorFor("ev3-ports:out" + shifterPort, "lego-ev3-m-motor")
	runner,		_ := ev3dev.TachoMotorFor("ev3-ports:out" + runnerPort, "lego-ev3-m-motor")
	shifter.SetPosition(0)
	runner.SetPosition(0)
	shifterr := Shifter {
		shifterMotor	: shifter,
		runnerMotor		: runner,
		offset			: offsetVal,
		rate			: shiftrt,
	}
	return shifterr
}

// id : range : [0, 3] (module number)
func (s Shifter) To(id int) {
	originalState, _ := s.shifterMotor.State()
	s.shifterMotor.SetStopAction("brake")
	s.shifterMotor.SetSpeedSetpoint(s.rate)
	s.shifterMotor.SetPositionSetpoint(id*s.offset).Command("run-to-abs-pos")
	for state, _ := s.shifterMotor.State(); state != originalState; state, _ = s.shifterMotor.State() {}
}

// target : degrees
// speed : motor units [unknown]
func (s Shifter) Run(target int, speed int) {
	originalState, _ := s.runnerMotor.State()
	s.runnerMotor.SetStopAction("brake")
	s.runnerMotor.SetSpeedSetpoint(speed)
	s.runnerMotor.SetPositionSetpoint(target).Command("run-to-rel-pos")
	for state, _ := s.runnerMotor.State(); state != originalState; state, _ = s.runnerMotor.State() {}
}

// target : degrees
// speed : motor units [unknown]
func (s Shifter) RunToAbs(target int, speed int) {
	originalState, _ := s.runnerMotor.State()
	s.runnerMotor.SetStopAction("brake")
	s.runnerMotor.SetSpeedSetpoint(speed)
	s.runnerMotor.SetPositionSetpoint(target).Command("run-to-abs-pos")
	for state, _ := s.runnerMotor.State(); state != originalState; state, _ = s.runnerMotor.State() {}
}
