package main

import (
	"strconv"
	"time"
)

const (
	P = float64(28)
	I = float64(10)
	D = float64(0)
	T = 20
)

func main() {
		r := InitRobot("A", "D", "2", "4", "1", "3")
		R := r
		R.colorLeft = r.colorMid
		s := InitShifter("B", "C", -1910, 800)
		l := InitLifter(s, 0, 1000, 0)
		c := InitColumn(s, 2)
		i := InitSlide(s, 3, 500)

		// green := r.ColorCalib("GREEN")
		white := r.ColorCalib("WHITE")
		black := r.ColorCalib("BLACK")
		grey := r.ColorCalib("GREY")
		greyM := R.ColorCalib("MGREY")
		pink := r.ColorCalib("PINK")
		// whiteM := r.ColorCalib("MWHITE")
		blackM := R.ColorCalib("MBLACK")
		AwaitButton()

		r.Move(400, black, T, P, I, D)
		r.Move(200, grey, T, P, I, D)
		R.Move(200, blackM, T, P, I, D)
		r.Follow(200, black, T)
		r.Follow(60, white, T)
		r.Move(-200, grey, T, P, I, D)
		r.Move(-200, white, T, P, I, D)
		r.Rotate(-24, 100)
		R.Move(200, greyM, T, P, I, D)
		R.Move(200, blackM, T, P, I, D)
		r.Run(80, 20)
		l.To(105)
		l.To(90)
		s.ToAsync(c.id)
		r.Move(-300, black, T, P, I, D)
		r.Rotate(121, 200)
		r.Follow(400, black, T)
		r.Follow(200, white, T)
		r.Follow(400, black, T)
		
		time.Sleep(time.Millisecond * 25)
		r.gyroSensor.SetMode("GYRO-RATE")
		time.Sleep(time.Millisecond * 25)
		r.gyroSensor.SetMode("GYRO-ANG")
		time.Sleep(time.Millisecond * 25)
		r.gyroSensor.SetMode("GYRO-ANG")

		c.To(960, 370)
		c.To(-960, 1000)
		r.Follow(200, white, T)
		r.Follow(400, black, T)
		r.Follow(200, white, T)
		r.Follow(300, pink, T)
		r.Follow(200, black, T)
		c.To(960, 1000)
		r.Move(-200, pink, T, P, I, D)
		c.To(-960, 1000)
		s.ToAsync(i.id)
		angs, _ := r.gyroSensor.Value(0)
		angi, _ := strconv.Atoi(angs)
		r.Rotate(-angi, 60)

		r.Move(-200, pink, T, P, I, D)
		r.Move(-300, white, T, P, I, D)
		r.Move(-200, grey, T, P, I, D)
		r.Move(-300, white, T, P, I, D)
		i.Unlock()
		s.ToAsync(0)
		r.Move(-300, grey, T, P, I, D)
		r.Move(-300, white, T, P, I, D)

		r.Move(-200, grey, T, P, I, D)
		r.Rotate(-30, 200)
		r.MoveTillButton(-250, P, I, D)
		l.AwaitTo()
}
