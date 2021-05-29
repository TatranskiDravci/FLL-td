package main

import (
	"github.com/ev3go/ev3dev"
	"math"
	"strconv"
	"time"
	"fmt"
)

type Robot struct {
	leftMotor		*ev3dev.TachoMotor
	rightMotor		*ev3dev.TachoMotor
	gyroSensor		*ev3dev.Sensor
	colorLeft		*ev3dev.Sensor
	colorMid		*ev3dev.Sensor
	colorRight		*ev3dev.Sensor
}

func InitRobot(leftPort, rightPort, gyroPort, clPort, cmPort, crPort string) Robot {
	left,	_	:= ev3dev.TachoMotorFor("ev3-ports:out" + leftPort, "lego-ev3-l-motor")
	right,	_	:= ev3dev.TachoMotorFor("ev3-ports:out" + rightPort, "lego-ev3-l-motor")
	gyro,	_	:= ev3dev.SensorFor("ev3-ports:in" + gyroPort, "lego-ev3-gyro")
	cl,	_	:= ev3dev.SensorFor("ev3-ports:in" + clPort, "lego-ev3-color")
	cm,		_	:= ev3dev.SensorFor("ev3-ports:in" + cmPort, "lego-ev3-color")
	cr,		_	:= ev3dev.SensorFor("ev3-ports:in" + crPort, "lego-ev3-color")

	time.Sleep(time.Millisecond * 50)
	gyro.SetMode("GYRO-RATE")
	gyro.SetMode("GYRO-ANG")
	return Robot {
		leftMotor	: left,
		rightMotor	: right,
		gyroSensor	: gyro,
		colorLeft	: cl,
		colorMid	: cm,
		colorRight	: cr,
	}
}

func (r Robot) Steering(speed int, u float64) {
	speedf := float64(speed)
	if speed > 0 {
		r.rightMotor.SetSpeedSetpoint(int(math.Max(math.Min(speedf + u, speedf), -speedf))).Command("run-forever")
		r.leftMotor.SetSpeedSetpoint(int(math.Max(math.Min(speedf - u, speedf), -speedf))).Command("run-forever")
	} else if speed < 0 {
		r.rightMotor.SetSpeedSetpoint(-int(math.Max(math.Min(speedf + u, speedf), -speedf))).Command("run-forever")
		r.leftMotor.SetSpeedSetpoint(-int(math.Max(math.Min(speedf - u, speedf), -speedf))).Command("run-forever")
	}
}

func (r Robot) Move(speed int, color [3]int, thresh int, P float64, I float64, D float64) {
	fmt.Println("call to Robot/Move")
	time.Sleep(50 * time.Millisecond)
	r.gyroSensor.SetMode("GYRO-RATE")
	r.gyroSensor.SetMode("GYRO-ANG")
	r.colorLeft.SetMode("RAW-RGB")
	time.Sleep(time.Millisecond * 50)

	prevAngi := 0
	tPrev := time.Now()
	tCol := time.Now()
	tStart := time.Now()
	E := float64(0)

	coloriR := 0
	coloriG := 0
	coloriB := 0

	r.leftMotor.SetStopAction("brake")
	r.rightMotor.SetStopAction("brake")

	r.Steering(speed, 0)

	for true {
		tNow := time.Now()
		angs, _ := r.gyroSensor.Value(0)
		angi, _ := strconv.Atoi(angs)

		if tNow.Sub(tCol).Milliseconds() >= 150 {
			colorsR, _ := r.colorLeft.Value(0)
			coloriR, _ = strconv.Atoi(colorsR)
			colorsG, _ := r.colorLeft.Value(1)
			coloriG, _ = strconv.Atoi(colorsG)
			colorsB, _ := r.colorLeft.Value(2)
			coloriB, _ = strconv.Atoi(colorsB)
			tCol = tNow
		}

		if Within(coloriR, color[0], thresh) && Within(coloriG, color[1], thresh) && Within(coloriB, color[2], thresh) {
			r.leftMotor.Command("stop")
			r.rightMotor.Command("stop")
			break
		}

		dt := tNow.Sub(tPrev).Seconds()
		E = E + LinInt(dt, float64(angi), float64(prevAngi))
		de := (float64(angi) - float64(prevAngi)) / dt
		u := P*float64(angi) + I*E + D*de

		fmt.Println(time.Now().Sub(tStart).Seconds(), angi, u, P*float64(angi), I*E, D*de)

		r.Steering(speed, u)
		prevAngi = angi
		tPrev = tNow
	}

	r.Rotate(0, 100)
	time.Sleep(time.Millisecond * 50)
	r.gyroSensor.SetMode("GYRO-RATE")
	r.gyroSensor.SetMode("GYRO-ANG")
}

func (r Robot) MoveTillButton(speed int, P float64, I float64, D float64) {
	fmt.Println("call to Robot/Move")
	time.Sleep(50 * time.Millisecond)
	r.gyroSensor.SetMode("GYRO-RATE")
	r.gyroSensor.SetMode("GYRO-ANG")
	time.Sleep(time.Millisecond * 50)

	prevAngi := 0
	tPrev := time.Now()
	E := float64(0)

	r.leftMotor.SetStopAction("brake")
	r.rightMotor.SetStopAction("brake")

	r.Steering(speed, 0)

	sig := false
	go ButtonSig(&sig)

	for !sig {
		tNow := time.Now()
		angs, _ := r.gyroSensor.Value(0)
		angi, _ := strconv.Atoi(angs)

		dt := tNow.Sub(tPrev).Seconds()
		E = E + LinInt(dt, float64(angi), float64(prevAngi))
		de := (float64(angi) - float64(prevAngi)) / dt
		u := P*float64(angi) + I*E + D*de

		r.Steering(speed, u)
		prevAngi = angi
		tPrev = tNow
	}
	r.leftMotor.Command("stop")
	r.rightMotor.Command("stop")
}

func (r Robot) Follow(speed int, color [3]int, thresh int) {
	r.colorLeft.SetMode("RGB-RAW")
	r.colorMid.SetMode("COL-COLOR")
	r.colorRight.SetMode("COL-COLOR")

	coloriR := 0
	coloriG := 0
	coloriB := 0
	tPrev := time.Now()
	for true {
		tNow := time.Now()
		if tNow.Sub(tPrev).Milliseconds() >= 150 {
			colorsR, _ := r.colorLeft.Value(0)
			coloriR, _ = strconv.Atoi(colorsR)
			colorsG, _ := r.colorLeft.Value(1)
			coloriG, _ = strconv.Atoi(colorsG)
			colorsB, _ := r.colorLeft.Value(2)
			coloriB, _ = strconv.Atoi(colorsB)
			tPrev = tNow
		}

		if Within(coloriR, color[0], thresh) && Within(coloriG, color[1], thresh) && Within(coloriB, color[2], thresh) {
			r.leftMotor.Command("stop")
			r.rightMotor.Command("stop")
			break
		}

		lcs, _ := r.colorMid.Value(0)
		rcs, _ := r.colorRight.Value(0)
		lci, _ := strconv.Atoi(lcs)
		rci, _ := strconv.Atoi(rcs)

		if lci == 1 && rci != 1 {
			r.rightMotor.SetSpeedSetpoint(int(1.25*float64(speed))).Command("run-forever")
			r.leftMotor.SetSpeedSetpoint(int(0.75*float64(speed))).Command("run-forever")
		} else if lci != 1 && rci == 1 {
			r.rightMotor.SetSpeedSetpoint(int(0.75*float64(speed))).Command("run-forever")
			r.leftMotor.SetSpeedSetpoint(int(1.25*float64(speed))).Command("run-forever")
		} else {
			r.rightMotor.SetSpeedSetpoint(speed).Command("run-forever")
			r.leftMotor.SetSpeedSetpoint(speed).Command("run-forever")
		}

	}
}

func (r Robot) Rotate(angle int, speed int) {
	r.leftMotor.SetStopAction("brake")
	r.rightMotor.SetStopAction("brake")
	for true {
		angs, _ := r.gyroSensor.Value(0)
		angi, _ := strconv.Atoi(angs)
		speedM := Modv(angle, angi, speed)
		if angi > angle {
			r.rightMotor.SetSpeedSetpoint(speedM).Command("run-forever")
			r.leftMotor.SetSpeedSetpoint(-speedM).Command("run-forever")
		} else if angi < angle {
			r.leftMotor.SetSpeedSetpoint(speedM).Command("run-forever")
			r.rightMotor.SetSpeedSetpoint(-speedM).Command("run-forever")
		} else if angi == angle {
			r.leftMotor.Command("stop")
			r.rightMotor.Command("stop")
			time.Sleep(100 * time.Millisecond)
			angs, _ := r.gyroSensor.Value(0)
			angi, _ := strconv.Atoi(angs)
			if angi == angle {
				break
			}
		}
	}
	time.Sleep(50 * time.Millisecond)
	r.gyroSensor.SetMode("GYRO-RATE")
	r.gyroSensor.SetMode("GYRO-ANG")
}

func (r Robot) ColorCalib(name string) [3]int {
	color, ok := GetColor(name)
	if !ok {
		r.colorLeft.SetMode("RGB-RAW")
		fmt.Println(name)
		AwaitButton()

		colorsR, _ := r.colorLeft.Value(0)
		colorsG, _ := r.colorLeft.Value(1)
		colorsB, _ := r.colorLeft.Value(2)

		R, _ := strconv.Atoi(colorsR)
		G, _ := strconv.Atoi(colorsG)
		B, _ := strconv.Atoi(colorsB)
		color = [3]int{R, G, B}

		SetColor(color, name)
	}
	return color
}

