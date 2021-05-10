package main

const black = 1
const P = float64(30)
const I = float64(100)

func main() {
	// initialize required modules
	r := InitRobot("A", "D", "2", "4")
	s := InitShifter("B", "C", -150, 200)
	c := InitCarrier(s, 1) 
	l := InitLift(s, 0)
	b := InitBox(s, 3)

	// path code
	r.MoveD(425, 250, 10, P, I)
	l.To(2.6, 420)
	r.Rotate(115, 100)
	r.Run(-300, 60)
	r.Rotate(-9, 100)
	r.MoveD(350, 80, 10, P, I)
	c.Release(300)

	l.ToUnbounded(-4.6, 400)
	r.MoveD(666, 200, 10, P, I)
	r.Rotate(-113, 100)

	r.MoveD(720, 200, 10, P, I)
	r.MoveD(725, 60, 5, P, I)
	l.To(6.9, 400)

	l.To(-5, 400)
	r.MoveD(580, 100, 10, P, I)
	r.Rotate(45, 100)
	r.Run(600, 80)

	r.Run(-400, 80)
	r.Rotate(52, 100)
	r.MoveD(130, 350, 20, P, I)
	b.Release(300)
}