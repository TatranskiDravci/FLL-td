package main

import (
	"fmt"
	"math"
	"time"
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
		speed = ModSpeed(target, angle, 45, max)

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
		speed = ModSpeed(target, angle, 45, max)

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
func (this Base) Move(max int, packet [2][3]int, sens Sensing, ctl PID) {
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

/*
	provides Move function, time interrupt
		max  		- max speed
		ctl	 		- PID parameters
		... 		- color sensing parameters
*/
func (this Base) MoveVar(max int, sens Sensing, ctl PID, breaking bool, packet ...[2][3]int) {
	fmt.Println("call to Robot/Move")
	this.ResetGyro()
	this.SetBrake()

	start := time.Now()
	speed := float64(max)
	this.RunSteering(speed, 0)
	ptime := 0.0
	perror := 0.0
	integral := 0.0
	i := 0
	color := packet[0]

	for {
		// breakpoint
		if sens.CompareColor(color) {
			this.Stop()
			if i < len(packet) {
				color = packet[i + 1]
				i++
				continue
			}
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
	if breaking {
		this.Stop()
		this.Rotate(-Measure(this.gyro, 0), 55)
	}
}

func (this Base) MoveEv(max int, sens Sensing, ctl PID, evpacket [2][3]int, evSP float64, packet ...[2][3]int) {
	this.MoveVar(max, sens, ctl, false, evpacket)
	ctl.SP = evSP
	this.MoveVar(max, sens, ctl, true, packet...)
}
