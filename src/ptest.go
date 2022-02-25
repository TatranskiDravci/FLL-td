package main

func main() {
	sens := NewSensing("3")
	pid  := NewPID(0.0, 25.0, 25.0, 0.0)
	base := NewBase("A", "D", "1")

	ev := sens.ColorCalib("green")
	p1 := sens.ColorCalib("black")

	base.MoveEv(300, sens, pid, ev, -90, p1)
}
