package main

const (
	T = 20
	P = 28
	I = 10
)

func main() {
	r := InitRobot("A", "D", "2", "4", "1", "3")
	pink := r.ColorCalib("")

	AwaitButton()
	for d := 0; d <= 6; d++ {
		for i := 0; i < 3; i++ {
			r.Move(400, pink, T, P, I, float64(d))
			AwaitButton()
		}
	}
}