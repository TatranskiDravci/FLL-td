package main

import (
	// "github.com/ev3go/ev3"
	"github.com/ev3go/ev3dev"
	// "github.com/ev3go/ev3dev/fb"
	"math"
	"strconv"
	"time"
	"fmt"
)

func modV(target int, current int, speed int) int{
	if current < (target - 10) || current > (target + 10) {
		return speed
	}
	return speed / 2
}

type Robot struct {
	leftMotor		*ev3dev.TachoMotor
	rightMotor		*ev3dev.TachoMotor
	gyroSensor		*ev3dev.Sensor
	colorSensor		*ev3dev.Sensor
	cl				*ev3dev.Sensor
	cr				*ev3dev.Sensor
}

func InitRobot(leftPort, rightPort, gyroPort, colorPort, clPort, crPort string) Robot {
	left,	_	:= ev3dev.TachoMotorFor("ev3-ports:out" + leftPort, "lego-ev3-l-motor")
	right,	_	:= ev3dev.TachoMotorFor("ev3-ports:out" + rightPort, "lego-ev3-l-motor")
	gyro,	_	:= ev3dev.SensorFor("ev3-ports:in" + gyroPort, "lego-ev3-gyro")
	color,	_	:= ev3dev.SensorFor("ev3-ports:in" + colorPort, "lego-ev3-color")
	cl,		_	:= ev3dev.SensorFor("ev3-ports:in" + clPort, "lego-ev3-color")
	cr,		_	:= ev3dev.SensorFor("ev3-ports:in" + crPort, "lego-ev3-color")

	time.Sleep(time.Millisecond * 50)
	gyro.SetMode("GYRO-RATE")
	gyro.SetMode("GYRO-ANG")
	robotr := Robot {
		leftMotor	: left,
		rightMotor	: right,
		gyroSensor	: gyro,
		colorSensor : color,
		cl			: cl,
		cr			: cr,
	}
	return robotr
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
	r.colorSensor.SetMode("RAW-RGB")
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
			colorsR, _ := r.colorSensor.Value(0)
			coloriR, _ = strconv.Atoi(colorsR)
			colorsG, _ := r.colorSensor.Value(1)
			coloriG, _ = strconv.Atoi(colorsG)
			colorsB, _ := r.colorSensor.Value(2)
			coloriB, _ = strconv.Atoi(colorsB)
			tCol = tNow
		}

		if within(coloriR, color[0], thresh) && within(coloriG, color[1], thresh) && within(coloriB, color[2], thresh) {
			r.leftMotor.Command("stop")
			r.rightMotor.Command("stop")
			break
		}

		dt := tNow.Sub(tPrev).Seconds()
		E = E + flinInt(dt, float64(angi), float64(prevAngi))
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
		E = E + flinInt(dt, float64(angi), float64(prevAngi))
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
	r.colorSensor.SetMode("RGB-RAW")
	r.cl.SetMode("COL-COLOR")
	r.cr.SetMode("COL-COLOR")

	coloriR := 0
	coloriG := 0
	coloriB := 0
	tPrev := time.Now()
	for true {
		tNow := time.Now()
		if tNow.Sub(tPrev).Milliseconds() >= 150 {
			colorsR, _ := r.colorSensor.Value(0)
			coloriR, _ = strconv.Atoi(colorsR)
			colorsG, _ := r.colorSensor.Value(1)
			coloriG, _ = strconv.Atoi(colorsG)
			colorsB, _ := r.colorSensor.Value(2)
			coloriB, _ = strconv.Atoi(colorsB)
			tPrev = tNow
		}

		if within(coloriR, color[0], thresh) && within(coloriG, color[1], thresh) && within(coloriB, color[2], thresh) {
			r.leftMotor.Command("stop")
			r.rightMotor.Command("stop")
			break
		}

		lcs, _ := r.cl.Value(0)
		rcs, _ := r.cr.Value(0)
		
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
		speedM := modV(angle, angi, speed)
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

func (r Robot) ColorCalib(color string) [3]int {
	r.colorSensor.SetMode("RGB-RAW")
	fmt.Println(color)
	AwaitButton()
	colorsR, _ := r.colorSensor.Value(0)
	R, _ := strconv.Atoi(colorsR)
	colorsG, _ := r.colorSensor.Value(1)
	G, _ := strconv.Atoi(colorsG)
	colorsB, _ := r.colorSensor.Value(2)
	B, _ := strconv.Atoi(colorsB)
	return [3]int{R, G, B}
}