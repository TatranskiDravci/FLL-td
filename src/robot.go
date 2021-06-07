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
	time.Sleep(time.Millisecond * 25)
	r.gyroSensor.SetMode("GYRO-RATE")
	r.colorLeft.SetMode("RGB-RAW")
	time.Sleep(time.Millisecond * 25)
	r.gyroSensor.SetMode("GYRO-ANG")
	time.Sleep(time.Millisecond * 25)
	r.gyroSensor.SetMode("GYRO-ANG")

	angiPrev := 0
	timeStart := time.Now()
	timePrev := time.Now()
	E := float64(0)

	coloriR := 0
	coloriG := 0
	coloriB := 0

	r.leftMotor.SetStopAction("brake")
	r.rightMotor.SetStopAction("brake")

	r.Steering(speed, 0)

	for true {
		timeNow := time.Now()
		angs, _ := r.gyroSensor.Value(0)
		angi, _ := strconv.Atoi(angs)

		colorsR, _ := r.colorLeft.Value(0)
		coloriR, _ = strconv.Atoi(colorsR)
		colorsG, _ := r.colorLeft.Value(1)
		coloriG, _ = strconv.Atoi(colorsG)
		colorsB, _ := r.colorLeft.Value(2)
		coloriB, _ = strconv.Atoi(colorsB)

		if Within(coloriR, color[0], thresh) && Within(coloriG, color[1], thresh) && Within(coloriB, color[2], thresh) {
			r.leftMotor.Command("stop")
			r.rightMotor.Command("stop")
			break
		}

		dt := timeNow.Sub(timePrev).Seconds()
		E = E + LinInt(dt, float64(angi), float64(angiPrev))
		de := (float64(angi) - float64(angiPrev)) / dt
		u := P*float64(angi) + I*E + D*de

		fmt.Println(timeNow.Sub(timeStart).Seconds(), angi, u, P*float64(angi), I*E, D*de)

		r.Steering(speed, u)
		angiPrev = angi
		timePrev = timeNow
	}

	angs, _ := r.gyroSensor.Value(0)
	angi, _ := strconv.Atoi(angs)
	if angi != 0 {
		r.Rotate(-angi, 60)
	}
}

func (r Robot) MoveTillButton(speed int, P float64, I float64, D float64) {
	time.Sleep(time.Millisecond * 25)
	r.gyroSensor.SetMode("GYRO-RATE")
	time.Sleep(time.Millisecond * 25)
	r.gyroSensor.SetMode("GYRO-ANG")
	time.Sleep(time.Millisecond * 25)
	r.gyroSensor.SetMode("GYRO-ANG")

	angiPrev := 0
	timePrev := time.Now()
	E := float64(0)

	r.leftMotor.SetStopAction("brake")
	r.rightMotor.SetStopAction("brake")

	r.Steering(speed, 0)

	sig := false
	go ButtonSig(&sig)

	for !sig {
		timeNow := time.Now()
		angs, _ := r.gyroSensor.Value(0)
		angi, _ := strconv.Atoi(angs)

		dt := timeNow.Sub(timePrev).Seconds()
		E = E + LinInt(dt, float64(angi), float64(angiPrev))
		de := (float64(angi) - float64(angiPrev)) / dt
		u := P*float64(angi) + I*E + D*de

		r.Steering(speed, u)
		angiPrev = angi
		timePrev = timeNow
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

	for true {
		colorsR, _ := r.colorLeft.Value(0)
		coloriR, _ = strconv.Atoi(colorsR)
		colorsG, _ := r.colorLeft.Value(1)
		coloriG, _ = strconv.Atoi(colorsG)
		colorsB, _ := r.colorLeft.Value(2)
		coloriB, _ = strconv.Atoi(colorsB)


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
			if speed >= 0 {
				r.rightMotor.SetSpeedSetpoint(int(1.25*float64(speed))).Command("run-forever")
				r.leftMotor.SetSpeedSetpoint(int(0.75*float64(speed))).Command("run-forever")
			} else {
				r.rightMotor.SetSpeedSetpoint(int(0.75*float64(speed))).Command("run-forever")
				r.leftMotor.SetSpeedSetpoint(int(1.25*float64(speed))).Command("run-forever")
			}
		} else if lci != 1 && rci == 1 {
			if speed >= 0 {
				r.rightMotor.SetSpeedSetpoint(int(0.75*float64(speed))).Command("run-forever")
				r.leftMotor.SetSpeedSetpoint(int(1.25*float64(speed))).Command("run-forever")
			} else {
				r.rightMotor.SetSpeedSetpoint(int(1.25*float64(speed))).Command("run-forever")
				r.leftMotor.SetSpeedSetpoint(int(0.75*float64(speed))).Command("run-forever")
			}
		} else {
			r.rightMotor.SetSpeedSetpoint(speed).Command("run-forever")
			r.leftMotor.SetSpeedSetpoint(speed).Command("run-forever")
		}

	}
}

func (r Robot) Rotate(angle int, speed int) {
	time.Sleep(time.Millisecond * 25)
	r.gyroSensor.SetMode("GYRO-RATE")
	time.Sleep(time.Millisecond * 25)
	r.gyroSensor.SetMode("GYRO-ANG")
	time.Sleep(time.Millisecond * 25)
	r.gyroSensor.SetMode("GYRO-ANG")
	

	r.leftMotor.SetStopAction("brake")
	r.rightMotor.SetStopAction("brake")
	minSpeed := float64(50)
	maxSpeed := float64(speed)
	for true {
		angs, _ := r.gyroSensor.Value(0)
		angi, _ := strconv.Atoi(angs)
		speedM := ModSpeed(float64(angle), float64(angi), minSpeed, maxSpeed)
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

func (r Robot) Run(speed, target int) {
	originalState, _ := r.leftMotor.State()
	r.leftMotor.SetStopAction("brake")
	r.leftMotor.SetSpeedSetpoint(speed)
	r.leftMotor.SetPositionSetpoint(target).Command("run-to-rel-pos")

	r.rightMotor.SetStopAction("brake")
	r.rightMotor.SetSpeedSetpoint(speed)
	r.rightMotor.SetPositionSetpoint(target).Command("run-to-rel-pos")

	for state, _ := r.leftMotor.State(); state != originalState; state, _ = r.leftMotor.State() {}
}