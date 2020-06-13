// Copyright 2016 Diego Elio Pettenò <flameeyes@flameeyes.com>
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
	"log"
)

var (
	profile = flag.Int("profile", 1, "Select profile 1 or 2 of the mouse.")
)

func main() {
	flag.Parse()

	if *profile < 1 || *profile > 2 {
		log.Fatalf("Invalid value for -profile: %v", *profile)
	}

	dev, err := device.Open()
	if err != nil {
		log.Fatal(err)
	}

	profileId := byte(*profile) - 1
	err = dev.SetProfile(profileId)
	if err != nil {
		log.Fatal(err)
	}
}
