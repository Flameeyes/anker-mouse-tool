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
	"fmt"
	"github.com/GeertJohan/go.hid"
	colorful "github.com/lucasb-eyer/go-colorful"
	"log"
	"strconv"
	"strings"
)

var (
	profile1Light = flag.String("profile1_light", "#0000ff:2:0", "String as color:brightness:breath for the light for profile #1.")
	profile2Light = flag.String("profile2_light", "#00ff00:2:0", "String as color:brightness:breath for the light for profile #2.")
)

const (
	HoltekVendorId     = 0x04d9
	AnkerMouseDeviceId = 0xfa50
)

func parseLightFlag(v string) (*colorful.Color, byte, byte, error) {
	p := strings.Split(v, ":")
	if len(p) != 3 {
		return nil, 0, 0, fmt.Errorf("Invalid profile light setting: %v", v)
	}

	c, err := colorful.Hex(p[0])
	if err != nil {
		return nil, 0, 0, err
	}

	bright, err := strconv.Atoi(p[1])
	if err != nil {
		return nil, 0, 0, err
	}

	breath, err := strconv.Atoi(p[2])
	if err != nil {
		return nil, 0, 0, err
	}

	return &c, byte(bright), byte(breath), nil
}

func main() {
	flag.Parse()

	c1, bright1, breath1, err := parseLightFlag(*profile1Light)
	if err != nil {
		log.Fatalf("Invalid value for -profile1_light: %v", err)
	}

	c2, bright2, breath2, err := parseLightFlag(*profile2Light)
	if err != nil {
		log.Fatalf("Invalid value for -profile2_light: %v", err)
	}

	device, err := hid.Open(HoltekVendorId, AnkerMouseDeviceId, "")
	if err != nil {
		log.Fatal(err)
	}

	cfg := NewConfig()
	cfg.Profiles[0].LightProfile.SetColor(*c1)
	cfg.Profiles[0].LightProfile.Brightness = bright1
	cfg.Profiles[0].LightProfile.BreathSpeed = breath1
	cfg.Profiles[1].LightProfile.SetColor(*c2)
	cfg.Profiles[1].LightProfile.Brightness = bright2
	cfg.Profiles[1].LightProfile.BreathSpeed = breath2

	err = cfg.Write(device)
	if err != nil {
		log.Fatal(err)
	}
}
