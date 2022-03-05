package main

import (
	"fmt"
	"math"
	"time"
)

const (
	_min = 20
	_initial = 0
	_final = 1
)

/*
	provides a Rotate method
		target - target  angle
		max    - max speed
*/
func (this Base) Rotate(target, max int) {
	this.ResetGyro()
	this.SetBrake()

	speed := max
	for {
		angle := Measure(this.gyro, 0)
		speed = ModSpeed(target, angle, _min, max)

		if angle == target {
			this.Stop()
			time.Sleep(100 * time.Millisecond)

			// "measure twice cut once"
			if Measure(this.gyro, 0) == target {
				return
			}
		} else if angle > target {
			this.RunTank( speed, -speed)
		} else {
			this.RunTank(-speed,  speed)
		}
	}
}


func (this Base) RotateR(target, max int) {
	this.ResetGyro()
	this.SetBrake()

	speed := max
	for {
		angle := Measure(this.gyro, 0)
		speed = ModSpeed(target, angle, _min, max)

		if angle == target {
			this.Stop()
			time.Sleep(100 * time.Millisecond)

			// "measure twice cut once"
			if Measure(this.gyro, 0) == target {
				return
			}
		} else if angle > target {
			this.RunTank( speed, 0)
		} else {
			this.RunTank(-speed, 0)
		}
	}
}

func (this Base) RotateL(target, max int) {
	this.ResetGyro()
	this.SetBrake()

	speed := max
	for {
		angle := Measure(this.gyro, 0)
		speed = ModSpeed(target, angle, _min, max)

		if angle == target {
			this.Stop()
			time.Sleep(100 * time.Millisecond)

			// "measure twice cut once"
			if Measure(this.gyro, 0) == target {
				return
			}
		} else if angle > target {
			this.RunTank(0, -speed)
		} else {
			this.RunTank(0,  speed)
		}
	}
}


/*
	provides speed modification for Rotate method
		target - target  angle
		angle  - current angle
		min    - min speed
		max    - max speed
*/
func ModSpeed(target, angle, min, max int) int {
	angleDelta := math.Abs(float64(target - angle))
	speedDelta := float64(max - min)

	return int(math.Min(0.005556 * speedDelta * angleDelta + float64(min), float64(max)))
}

/*
	provides Move function, time interrupt
		max  		- max speed
		ctl	 		- PID parameters
		... 		- color sensing parameters
*/
func (this Base) Move(max int, sens Sensing, ctl PID, packet [2][3]int) {
	fmt.Println("call to Robot/Move")
	this.ResetGyro()
	this.SetBrake()

	start := time.Now()
	speed := float64(max)
	this.RunSteering(speed, 0)
	ptime := 0.0
	perror := 0.0
	integral := 0.0

	for {
		// breakpoint
		if sens.CompareColor(packet) {
			this.Stop()
			break
		}

		// movement correction
		angle := Measure(this.gyro, 0)
		time := time.Now().Sub(start).Seconds()

		error := ctl.SP - float64(angle)

		dtime := time - ptime
		ptime = time

		integral = integral + 0.5 * dtime * (perror + error)
		// derror := error - perror
		perror = error


	 	steering := ctl.K_P * error + ctl.K_I * integral // + ctl.K_D * derror / dtime
		this.RunSteering(speed, -steering)
	}
	this.Rotate(-Measure(this.gyro, 0), 55)
}


func (this Base) MoveTimed(max int, target float64, ctl PID) {
	fmt.Println("call to Robot/Move")
	this.ResetGyro()
	this.SetBrake()

	start := time.Now()
	speed := float64(max)
	this.RunSteering(speed, 0)
	ptime := 0.0
	perror := 0.0
	integral := 0.0

	for {
		// movement correction
		angle := Measure(this.gyro, 0)
		time := time.Now().Sub(start).Seconds()

		error := ctl.SP - float64(angle)

		dtime := time - ptime
		ptime = time

		integral = integral + 0.5 * dtime * (perror + error)
		// derror := error - perror
		perror = error


	 	steering := ctl.K_P * error + ctl.K_I * integral // + ctl.K_D * derror / dtime
		this.RunSteering(speed, -steering)

		// breakpoint
		if time >= target {
			break
		}
	}
	this.Stop()
	this.Rotate(-Measure(this.gyro, 0), 70)
}

func (this Base) MoveTimedReverse(max int, target float64, ctl PID) {
	fmt.Println("call to Robot/Move")
	this.ResetGyro()
	this.SetBrake()

	start := time.Now()
	speed := float64(max)
	this.RunSteeringReverse(speed, 0)
	ptime := 0.0
	perror := 0.0
	integral := 0.0

	for {
		// movement correction
		angle := Measure(this.gyro, 0)
		time := time.Now().Sub(start).Seconds()

		error := ctl.SP - float64(angle)

		dtime := time - ptime
		ptime = time

		integral = integral + 0.5 * dtime * (perror + error)
		// derror := error - perror
		perror = error


	 	steering := ctl.K_P * error + ctl.K_I * integral // + ctl.K_D * derror / dtime
		this.RunSteeringReverse(speed, steering)

		// breakpoint
		if time >= target {
			break
		}
	}
	this.Stop()
	this.Rotate(-Measure(this.gyro, 0), 70)
}

func (this Base) MoveReverse(max int, sens Sensing, ctl PID, packet [2][3]int) {
	fmt.Println("call to Robot/Move")

	this.ResetGyro()
	this.SetBrake()

	start := time.Now()
	speed := float64(max)
	this.RunSteering(speed, 0)
	ptime := 0.0
	perror := 0.0
	integral := 0.0

	for {
		// breakpoint
		if sens.CompareColor(packet) {
			break
		}

		// movement correction
		angle := Measure(this.gyro, 0)
		time := time.Now().Sub(start).Seconds()

		error := ctl.SP - float64(angle)

		dtime := time - ptime
		ptime = time

		integral = integral + 0.5 * dtime * (perror + error)
		// derror := error - perror
		perror = error


	 	steering := ctl.K_P * error + ctl.K_I * integral // + ctl.K_D * derror / dtime
		this.RunSteering(speed, steering)
	}
	this.Stop()
	this.Rotate(-Measure(this.gyro, 0), 55)
}

/*
	provides non-stop Move function, time interrupt
		max  		- max speed
		ctl	 		- PID parameters
		... 		- color sensing parameters
*/
func (this Base) MoveNS(max int, sens Sensing, ctl *PIDNS, packet [2][3]int, kind int) {
	fmt.Println("call to Robot/Move")
	this.SetBrake()

	if kind == _initial {
		this.ResetGyro()
		ctl.start = time.Now()
		ctl.ptime = 0.0
		ctl.perror = 0.0
		ctl.integral = 0.0
	}



	speed := float64(max)
	this.RunSteering(speed, 0)
	start := ctl.start
	ptime := ctl.ptime
	perror := ctl.perror
	integral := ctl.integral

	for {
		// breakpoint
		if sens.CompareColor(packet) {
			break
		}

		// movement correction
		angle := Measure(this.gyro, 0)
		time := time.Now().Sub(start).Seconds()

		error := ctl.SP - float64(angle)

		dtime := time - ptime
		ptime = time

		integral = integral + 0.5 * dtime * (perror + error)
		// derror := error - perror
		perror = error


	 	steering := ctl.K_P * error + ctl.K_I * integral // + ctl.K_D * derror / dtime
		this.RunSteering(speed, -steering)
	}

	ctl.ptime = ptime
	ctl.perror = perror
	ctl.integral = integral

	if kind == _final {
		this.Stop()
		this.Rotate(-Measure(this.gyro, 0), 55)
	}
}
