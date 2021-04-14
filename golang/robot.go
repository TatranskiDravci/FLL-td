package main

import (
	// "github.com/ev3go/ev3"
	"github.com/ev3go/ev3dev"
	// "github.com/ev3go/ev3dev/fb"
	"math"
	"fmt"
	"strconv"
	"time"
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
	ultraSensor		*ev3dev.Sensor
}

func InitRobot(leftPort string, rightPort string, gyroPort string, ultraPort string) Robot {
	left,	_	:= ev3dev.TachoMotorFor("ev3-ports:out" + leftPort, "lego-ev3-l-motor")
	right,	_	:= ev3dev.TachoMotorFor("ev3-ports:out" + rightPort, "lego-ev3-l-motor")
	gyro,	_	:= ev3dev.SensorFor("ev3-ports:in" + gyroPort, "lego-ev3-gyro")
	ultra,	_	:= ev3dev.SensorFor("ev3-ports:in" + ultraPort, "lego-ev3-us") 
	robot := Robot {
		leftMotor	: left,
		rightMotor	: right,
		gyroSensor	: gyro,
		ultraSensor : ultra,
	}
	return robot
}

func (r Robot) Steering(speed int, direction int, turnat int) {
	speedf := float64(speed)
	directionf := float64(direction)
	turnatf := 1.0/float64(turnat)
	if speed > 0 {
		if direction >= 0 {
			r.leftMotor.SetSpeedSetpoint(int(math.Max(math.Min(turnatf*speedf*directionf + speedf, speedf), -speedf))).
				Command("run-forever")
			r.rightMotor.SetSpeedSetpoint(int(math.Max(math.Min(speedf - turnatf*speedf*directionf, speedf), -speedf))).
				Command("run-forever")
		} else {
			r.rightMotor.SetSpeedSetpoint(int(math.Max(math.Min(speedf - turnatf*speedf*directionf, speedf), -speedf))).
				Command("run-forever")
			r.leftMotor.SetSpeedSetpoint(int(math.Max(math.Min(turnatf*speedf*directionf + speedf, speedf), -speedf))).
				Command("run-forever")
		}
	} else if speed < 0 {
		if direction >= 0 {
			r.leftMotor.SetSpeedSetpoint(-int(math.Max(math.Min(turnatf*speedf*directionf + speedf, speedf), -speedf))).
				Command("run-forever")
			r.rightMotor.SetSpeedSetpoint(-int(math.Max(math.Min(speedf - turnatf*speedf*directionf, speedf), -speedf))).
				Command("run-forever")
		} else {
			r.rightMotor.SetSpeedSetpoint(-int(math.Max(math.Min(speedf - turnatf*speedf*directionf, speedf), -speedf))).
				Command("run-forever")
			r.leftMotor.SetSpeedSetpoint(-int(math.Max(math.Min(turnatf*speedf*directionf + speedf, speedf), -speedf))).
				Command("run-forever")
		}
	}
}

func (r Robot) Move(speed int, distance int, threshold int, turnat int) {
	r.leftMotor.SetStopAction("brake")
	r.rightMotor.SetStopAction("brake")
	r.gyroSensor.SetMode("GYRO-RATE")
	r.gyroSensor.SetMode("GYRO-ANG")
	r.Steering(speed, 0, turnat)
	prevAngi := 0
	for true {
		angs,	_ := r.gyroSensor.Value(0)
		angi,	_ := strconv.Atoi(angs)
		dists,	_ := r.ultraSensor.Value(0)
		disti,	_ := strconv.Atoi(dists)
		fmt.Println(angi, disti)
		if disti < (distance + threshold) && disti > (distance - threshold) {
			r.leftMotor.Command("stop")
			r.rightMotor.Command("stop")
			time.Sleep(50 * time.Millisecond)
			dists, _ := r.ultraSensor.Value(0)
			disti, _ := strconv.Atoi(dists)
			if disti < (distance + threshold) && disti > (distance - threshold) {
				break
			}
		} else if disti < distance {
			if angi != prevAngi {
				r.Steering(speed, -angi, turnat)
			}
		} else if  disti > distance {
			if angi != prevAngi {
				r.Steering(-speed, -angi, turnat)
			}
		}
		prevAngi = angi
	}
}

func (r Robot) Rotate(speed int, angle int) {
	r.leftMotor.SetStopAction("brake")
	r.rightMotor.SetStopAction("brake")
	r.gyroSensor.SetMode("GYRO-RATE")
	r.gyroSensor.SetMode("GYRO-ANG")
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
			time.Sleep(50 * time.Millisecond)
			angs, _ := r.gyroSensor.Value(0)
			angi, _ := strconv.Atoi(angs)
			if angi == angle {
				break
			}
		}
	}
}
