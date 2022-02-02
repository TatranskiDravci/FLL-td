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

	white := r.ColorCalib("WHITE")
	black := r.ColorCalib("BLACK")
	green := r.ColorCalib("GREEN")
	pink := r.ColorCalib("PINK")
	blackM := R.ColorCalib("MBLACK")
	// whiteM := R.ColorCalib("MWHITE")
	pinkM := R.ColorCalib("MPINK")
	yellow := r.ColorCalib("YELLOW")
	
	AwaitButton()
	

	R.Move(600, blackM, T, P, I, D)
	r.Follow(400, green, T)
	r.FollowUntilNot(400, green, T)
	r.Follow(400, green, T)
	r.Follow(300, black, T)
	l.To(0)
	s.ToAsync(c.id)
	r.Follow(300, green, T)
	r.Follow(400, black, T)
	c.To(900, 200)
	c.To(-900, 1000)
	r.Follow(400, pink, T)
	s.ToAsync(l.id)
	r.Follow(200, black, T)
	r.Rotate(-40, 100)
	l.To(15)
	r.Move(200, yellow, T, P, I, D)
	l.To(0)
	r.Rotate(-40, 150)
	r.Run(60, 60)
	l.To(15)
	R.Move(200, pinkM, T, P, I, D)
	r.Rotate(80, 200)
	r.Move(400, green, T, P, I, D)
	r.Move(300, white, T, P, I, D)
	r.Move(400, green, T, P, I, D)
	r.Move(300, white, T, P, I, D)
	r.Rotate(90, 200)
	r.MoveTillButton(300, P, I, D)
	// s.ToAsync(0)
	// r.Rotate(78, 200)
	// R.Move(-300, blackM, T, P, I, D)
	// R.Move(-300, whiteM, T, P, I, D)
	// r.MoveTillButton(-200, P, I, D)
	// s.AwaitTo()
}