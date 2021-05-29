package main

import "github.com/ev3go/ev3dev"

type Shifter struct {
	shifterMotor	*ev3dev.TachoMotor
	runnerMotor		*ev3dev.TachoMotor
	offset			int
	rate			int
	startState		ev3dev.MotorState
	current			int
}

func InitShifter(shifterPort string, runnerPort string, offsetVal int, shiftrt int) Shifter {
	shifter,	_ := ev3dev.TachoMotorFor("ev3-ports:out" + shifterPort, "lego-ev3-m-motor")
	runner,		_ := ev3dev.TachoMotorFor("ev3-ports:out" + runnerPort, "lego-ev3-l-motor")
	shifter.SetPosition(0)
	runner.SetPosition(0)
	start, _ := shifter.State()
	return Shifter {
		shifterMotor	: shifter,
		runnerMotor		: runner,
		offset			: offsetVal,
		rate			: shiftrt,
		startState		: start,
		current			: 0,
	}
}

func (s *Shifter) To(id int) {
	s.current = id
	originalState, _ := s.shifterMotor.State()
	s.shifterMotor.SetStopAction("brake")
	s.shifterMotor.SetSpeedSetpoint(s.rate)
	s.shifterMotor.SetPositionSetpoint(id*s.offset).Command("run-to-abs-pos")
	for state, _ := s.shifterMotor.State(); state != originalState; state, _ = s.shifterMotor.State() {}
}

func (s *Shifter) ToAsync(id int) {
	s.current = id
	originalState, _ := s.shifterMotor.State()
	s.shifterMotor.SetStopAction("brake")
	s.shifterMotor.SetSpeedSetpoint(s.rate)
	s.shifterMotor.SetPositionSetpoint(id*s.offset).Command("run-to-abs-pos")
	for state, _ := s.shifterMotor.State(); state != originalState; state, _ = s.shifterMotor.State() {}
}

func (s Shifter) Run(target int, speed int) {
	originalState, _ := s.runnerMotor.State()
	s.runnerMotor.SetStopAction("brake")
	s.runnerMotor.SetSpeedSetpoint(speed)
	s.runnerMotor.SetPositionSetpoint(target).Command("run-to-rel-pos")
	for state, _ := s.runnerMotor.State(); state != originalState; state, _ = s.runnerMotor.State() {}
}

func (s Shifter) RunAsync(target int, speed int) {
	s.runnerMotor.SetStopAction("brake")
	s.runnerMotor.SetSpeedSetpoint(speed)
	s.runnerMotor.SetPositionSetpoint(target).Command("run-to-rel-pos")
}

func (s Shifter) AwaitTo() {
	for state, _ := s.shifterMotor.State(); state != s.startState; state, _ = s.runnerMotor.State() {}
}

func (s Shifter) AwaitRun() {
	for state, _ := s.runnerMotor.State(); state != s.startState; state, _ = s.runnerMotor.State() {}
}