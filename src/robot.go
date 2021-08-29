package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/ev3go/ev3dev"
)

type Robot struct {
	leftMotor  *ev3dev.TachoMotor
	rightMotor *ev3dev.TachoMotor
	gyroSensor *ev3dev.Sensor
	colorLeft  *ev3dev.Sensor
	colorMid   *ev3dev.Sensor
	colorRight *ev3dev.Sensor
	angle      int
}

func InitRobot(leftPort, rightPort, gyroPort, clPort, cmPort, crPort string) Robot {
	left, _ := ev3dev.TachoMotorFor("ev3-ports:out"+leftPort, "lego-ev3-l-motor")
	right, _ := ev3dev.TachoMotorFor("ev3-ports:out"+rightPort, "lego-ev3-l-motor")
	gyro, _ := ev3dev.SensorFor("ev3-ports:in"+gyroPort, "lego-ev3-gyro")
	cl, _ := ev3dev.SensorFor("ev3-ports:in"+clPort, "lego-ev3-color")
	cm, _ := ev3dev.SensorFor("ev3-ports:in"+cmPort, "lego-ev3-color")
	cr, _ := ev3dev.SensorFor("ev3-ports:in"+crPort, "lego-ev3-color")

	time.Sleep(time.Millisecond * 50)
	gyro.SetMode("GYRO-RATE")
	gyro.SetMode("GYRO-ANG")
	return Robot{
		leftMotor:  left,
		rightMotor: right,
		gyroSensor: gyro,
		colorLeft:  cl,
		colorMid:   cm,
		colorRight: cr,
		angle:      0,
	}
}

func (r Robot) Rotate(angle int, speed int) {
	r.gyroSensor.SetMode("GYRO-ANG")

	r.leftMotor.SetStopAction("brake")
	r.rightMotor.SetStopAction("brake")
	minSpeed := float64(50)
	maxSpeed := float64(speed)
	for {
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
