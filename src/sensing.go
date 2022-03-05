package main

import (
	"github.com/ev3go/ev3dev"
	"strconv"
	"math"
	"os/exec"
	"os"
	"io/ioutil"
	"io/fs"
	"fmt"
)

/*
	cl - left   color sensor
*/
type Sensing struct {
	c *ev3dev.Sensor
	k [3]float64
	l [3]int
}

/*
	p - left color sensor port
*/
func NewSensing(p string) Sensing {
	c, _ := ev3dev.SensorFor("ev3-ports:in" + p, "lego-ev3-color")

	c.SetMode("RGB-RAW")

	return Sensing {
		c: c,
		k: [3]float64{1.,1.,1.},	// default profile
		l: [3]int{0.,0.,0.},		// default profile
	}
}

/*
	provides a sensor measurement function
		sensor  - a sensor to use for the measurement
		address - measurement address
*/
func Measure(sensor *ev3dev.Sensor, address int) int {
	mString, _ := sensor.Value(address)
	m, _ := strconv.Atoi(mString)
	return m
}

// func MeasureString(sensor *ev3dev.Sensor, address int) string {
// 	mString, _ := sensor.Value(address)
// 	return mString
// }


// provides a sensor measurement function on Sensing
func (this Sensing) Measure() [3]int {
	R := int(this.k[0] * float64(Measure(this.c, 0) - this.l[0]))
	G := int(this.k[1] * float64(Measure(this.c, 1) - this.l[1]))
	B := int(this.k[2] * float64(Measure(this.c, 2) - this.l[2]))
	return [3]int {R, G, B}
}

func (this Sensing) MeasureString() [3]string {
	out := this.Measure()
	return [3]string{strconv.Itoa(out[0]), strconv.Itoa(out[1]), strconv.Itoa(out[2])}
}

func (this Sensing) CompareColor(packet [2][3]int) bool {
	mR, _ := this.c.Value(0)
	mG, _ := this.c.Value(1)
	mB, _ := this.c.Value(2)

	R, _ := strconv.Atoi(mR)
	G, _ := strconv.Atoi(mG)
	B, _ := strconv.Atoi(mB)

	dR := int(math.Abs(float64(R - packet[0][0])))
	dG := int(math.Abs(float64(G - packet[0][1])))
	dB := int(math.Abs(float64(B - packet[0][2])))

	if dR <= 50 || dG <= 50 || dB <= 50 {
		fmt.Println(R, G, B, dR, dG, dB)
	}

	return dR <= packet[1][0] && dG <= packet[1][1] && dB <= packet[1][2]
}

func (this Sensing) ColorCalib(name string) [2][3]int {
	packet, ok := GetColor2(name)
	if !ok {
		fmt.Println("COLOR " + name + " NOT SET")
		AwaitButton()
		fmt.Println("<<<CALIBRATING>>>")
		out := ""
		for i := 0; i < 500; i++ {
			co := this.MeasureString()
			out = out + co[0] + " " + co[1] + " " + co[2] + "\n"
		}
		_, err := ioutil.ReadFile("../data/temp")
		if err != nil {
			os.Create("../data/temp")
		}
		ioutil.WriteFile("../data/temp", []byte(out), fs.ModePerm)

		cmd := exec.Command("./calib", name)
		cmd.Run()
	}


	return packet
}

func (this Sensing) ProfileCalib(id string) {
	k, l, ok := GetProfile(id)

	if !ok {
		fmt.Println("MISSING PROFILE " + id)

		fmt.Println("SCAN BLACK POINT")
		AwaitButton()
		fmt.Println("<<<CALIBRATING>>>")

		L := this.Measure()

		fmt.Println("SCAN WHITE POINT")
		AwaitButton()
		fmt.Println("<<<CALIBRATING>>>")

		H := this.Measure()

		this.k = [3]float64{255./float64(H[0] - L[0]), 255./float64(H[1] - L[1]), 255./float64(H[2] - L[2])}
		this.l = L

		SetProfile(this.k, this.l, id)

		fmt.Println("PROFILE " + id + " LOADED AND SAVED")
		return
	}

	this.k = k
	this.l = l
	fmt.Println("PROFILE " + id + " LOADED")
}
