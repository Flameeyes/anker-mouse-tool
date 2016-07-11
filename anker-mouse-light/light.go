// Copyright 2016 Diego Elio Pettenò <flameeyes@flameeyes.eu>
//
// Permission is hereby granted, free of charge, to any person
// obtaining a copy of this software and associated documentation
// files (the "Software"), to deal in the Software without
// restriction, including without limitation the rights to use, copy,
// modify, merge, publish, distribute, sublicense, and/or sell copies
// of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS
// BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN
// ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
// CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package main

import (
	"flag"
	"github.com/flameeyes/anker-mouse-tool/device"
	colorful "github.com/lucasb-eyer/go-colorful"
	"log"
)

var (
	lightColor  = flag.String("light_color", "#0000ff", "Colour to set the light to.")
	brightness  = flag.Int("brightness", 2, "Brightness of the device light, between 0 and 3 (0 means off.")
	breathSpeed = flag.Int("breath_speed", 0, "Speed of the \"breath\" of the device light, between 0 and 3 (0 means the light stays always-on.)")
)

type SetLightReport struct {
	ReportId     byte // 0x02
	InternalId   byte // 0x04
	InverseRed   byte
	InverseGreen byte
	InverseBlue  byte
	Brightness   byte
	BreathSpeed  byte
	Unknown      [9]byte // All zeroes.
}

func main() {
	flag.Parse()

	if *brightness < 0 || *brightness > 3 {
		log.Fatalf("Invalid value for -brightness: %v", *brightness)
	}

	if *breathSpeed < 0 || *breathSpeed > 3 {
		log.Fatalf("Invalid value for -breath_speed: %v", *breathSpeed)
	}

	c, err := colorful.Hex(*lightColor)
	if err != nil {
		log.Fatal(err)
	}

	dev, err := device.Open()
	if err != nil {
		log.Fatal(err)
	}

	r, g, b := c.RGB255()

	report := SetLightReport{
		ReportId:     0x02,
		InternalId:   0x04,
		InverseRed:   ^r,
		InverseGreen: ^g,
		InverseBlue:  ^b,
		Brightness:   byte(*brightness),
		BreathSpeed:  byte(*breathSpeed),
	}

	err = dev.WriteFeatureReport(report)
	if err != nil {
		log.Fatal(err)
	}
}
