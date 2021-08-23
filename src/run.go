package main

const (
	P = float64(26)
	I = float64(8)
	D = float64(0)
	T = 20
)

func main() {
		r := InitRobot("A", "D", "2", "4", "1", "3")
		s := InitShifter("B", "C", -1910, 800)
		d := InitDropper(s, 1, 500)

		green := r.ColorCalib("SGREEN")
		AwaitButton()

		s.ToAsync(d.id)
		r.Move(250, green, T, P, I, D)
		d.Release()
		s.ToAsync(0)
		r.MoveTillButton(-500, P, I, D)
		s.AwaitTo()
}