package main

const (
	P = float64(28)
	I = float64(10)
	D = float64(2)
	T = 20
)

func main() {
	r := InitRobot("A", "D", "2", "4", "1", "3")
	R := r
	R.colorLeft = r.colorMid
	s := InitShifter("B", "C", -1910, 800)
	l := InitLifter(s, 0, 1000, 100)
	c := InitColumn(s, 2)
	e := InitEuler(s, 4, 500)

	// white := r.ColorCalib("WHITE")
	black := r.ColorCalib("BLACK")
	green := r.ColorCalib("GREEN")
	pink := r.ColorCalib("PINK")
	blackM := R.ColorCalib("MBLACK")
	whiteM := R.ColorCalib("MWHITE")
	yellow := r.ColorCalib("YELLOW")
	
	AwaitButton()
	

	R.Move(400, blackM, T, P, I, D)
	r.Follow(400, black, T)
	r.Follow(200, green, T)
	r.Follow(200, black, T)
	l.To(0)
	s.ToAsync(c.id)
	r.Follow(200, green, T)
	r.Follow(200, black, T)
	c.To(-360, 100)
	c.To(500, 600)
	s.ToAsync(e.id)
	r.Follow(200, pink, T)
	r.Run(100, -150)
	e.Down()
	r.Follow(200, black, T)
	r.Move(-200, green, T, P, I, D)
	e.Up()
	s.ToAsync(l.id)
	r.Follow(200, black, T)
	r.Rotate(-38, 100)
	l.To(15)
	r.Move(200, yellow, T, P, I, D)
	l.To(0)
	r.Rotate(-40, 150)
	l.To(15)
	s.ToAsync(0)
	r.Rotate(78, 200)
	R.Move(-300, blackM, T, P, I, D)
	R.Move(-300, whiteM, T, P, I, D)
	r.MoveTillButton(-200, P, I, D)
	s.AwaitTo()
}