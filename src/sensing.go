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
	c - color sensor
*/
type Sensing struct {
	c *ev3dev.Sensor
}

/*
	p - color sensor port
*/
func NewSensing(p string) Sensing {
	c, _ := ev3dev.SensorFor("ev3-ports:in" + p, "lego-ev3-color")

	c.SetMode("RGB-RAW")

	return Sensing {
		c: c,
	}
}

/*
	provides a sensor reading function
		sensor  - a sensor to use for the measurement
		address - measurement address
*/
func Measure(sensor *ev3dev.Sensor, address int) int {
	mString, _ := sensor.Value(address)
	m, _ := strconv.Atoi(mString)
	return m
}


/*
	provides a color reading function on Sensing
*/
func (this Sensing) Measure() [3]int {
	return [3]int {Measure(this.c, 0), Measure(this.c, 1), Measure(this.c, 2)}
}

/*
	provides a color string reading function on Sensing
*/
func (this Sensing) MeasureString() [3]string {
	out := this.Measure()
	return [3]string{strconv.Itoa(out[0]), strconv.Itoa(out[1]), strconv.Itoa(out[2])}
}

/*
	provides a color reading & comparison function on Sensing
		packet	- color packet to be compared
*/
func (this Sensing) CompareColor(packet [2][3]int) bool {
	// read sensor directly to improve performance and decrease overhead
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

/*
	provides a color calibration function on Sensing
		name	- color name (used when creating entry in data/)
*/
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
