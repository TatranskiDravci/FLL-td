package main

const (
	P = float64(28)
	I = float64(10)
	D = float64(0)
	T = 20
)

func main() {
		r := InitRobot("A", "D", "2", "4", "1", "3")
		s := InitShifter("B", "C", -1910, 800)
		d := InitDropper(s, 1, 500)
		green := r.ColorCalib("green")
		
		AwaitButton()
		s.ToAsync(d.id)

		r.Move(400, green, T, P, I, D)
		s.AwaitTo()
		d.Release()
		r.MoveTillButton(-400, P, I, D)
		s.To(0)
}