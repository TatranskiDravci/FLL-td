package main

import (
	// "github.com/ev3go/ev3"
	"github.com/ev3go/ev3dev"
	// "github.com/ev3go/ev3dev/fb"
	"math"
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

	time.Sleep(time.Millisecond * 50)
	gyro.SetMode("GYRO-RATE")
	gyro.SetMode("GYRO-ANG")
	robotr := Robot {
		leftMotor	: left,
		rightMotor	: right,
		gyroSensor	: gyro,
		ultraSensor : ultra,
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


func (r Robot) MoveD(distance int, speed int, threshold int, P float64, I float64) {
	time.Sleep(time.Millisecond * 50)

	distLower := distance - threshold
	distUpper := distance + threshold
	prevAngi := 0
	tPrev := time.Now()
	tDist1 := time.Now()
	tDist2 := time.Now()
	E := float64(0)

	r.leftMotor.SetStopAction("brake")
	r.rightMotor.SetStopAction("brake")

	r.Steering(speed, 0)

	dists, _ := r.ultraSensor.Value(0)
	disti, _ := strconv.Atoi(dists)
	for true {
		tNow := time.Now()
		tDist2 = time.Now()
		dtDist := tDist2.Sub(tDist1).Milliseconds() 

		angs, _ := r.gyroSensor.Value(0)
		angi, _ := strconv.Atoi(angs)
		if dtDist >= 250 {
			dists,	_ = r.ultraSensor.Value(0)
			disti,	_ = strconv.Atoi(dists)
			tDist1 = time.Now()
		}

		E = E + flinInt(tNow.Sub(tPrev).Seconds(), float64(angi), float64(prevAngi))
		/* de := (float64(angi) - float64(prevAngi)) / dt */

		u := P*float64(angi) + I*E /* + PID[2]*de */

		if disti < distance {
			r.Steering(modV(distance, disti, speed), u)
		} else if disti > distance {
			r.Steering(modV(distance, disti, -speed), u)
		}
		 
		if disti <= distUpper && disti >= distLower {
			r.leftMotor.Command("stop")
			r.rightMotor.Command("stop")
			time.Sleep(100 * time.Millisecond)
			dists, _ := r.ultraSensor.Value(0)
			disti, _ := strconv.Atoi(dists)
			if disti < distUpper && disti > distLower {
				break
			}
		}
		prevAngi = angi
		tPrev = tNow
	}

	r.Rotate(0, 100)
	time.Sleep(time.Millisecond * 50)
	r.gyroSensor.SetMode("GYRO-RATE")
	r.gyroSensor.SetMode("GYRO-ANG")
}

// angle : degrees (measured by the gyro sensor)
// speed : motor units [unknown ]
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

func (r Robot) Run(target int, speed int) {
	originalState, _ := r.leftMotor.State()
	r.leftMotor.SetStopAction("brake")
	r.rightMotor.SetStopAction("brake")
	r.leftMotor.SetSpeedSetpoint(speed)
	r.rightMotor.SetSpeedSetpoint(speed)
	r.leftMotor.SetPositionSetpoint(target).Command("run-to-rel-pos")
	r.rightMotor.SetPositionSetpoint(target).Command("run-to-rel-pos")
	for true {
		leftState,  _ := r.leftMotor.State()
		rightState, _ := r.rightMotor.State()
		if leftState == originalState {
			r.leftMotor.Command("stop")
		}
		if rightState == originalState {
			r.rightMotor.Command("stop")
		}
		if leftState == originalState && rightState == originalState {
			break
		}
	}
}
