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
	"bytes"
	"encoding/binary"
	"flag"
	"github.com/GeertJohan/go.hid"
	"log"
)

var (
	profile = flag.Int("profile", 1, "Select profile 1 or 2 of the mouse.")
)

const (
	HoltekVendorId     = 0x04d9
	AnkerMouseDeviceId = 0xfa50
)

type SetProfileReport1 struct {
	ReportId  byte    // 0x02
	Constant1 [7]byte // 0x02, 0x40, 0x00, 0x01, 0x00, 0xFA, 0xFA
	Profile   byte
	Unknown   [7]byte // All zeroes.
}

var SetProfile1Constant1 = [7]byte{0x02, 0x40, 0x00, 0x01, 0x00, 0xFA, 0xFA}

type SetProfileReport2 struct {
	ReportId  byte   // 0x02
	Constant1 uint16 // 0x0101
	Profile   byte
	Unknown   [12]byte // All zeroes.
}

func WriteFeatureReport(device *hid.Device, report interface{}) error {

	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, report)
	if err != nil {
		return err
	}

	_, err = device.SendFeatureReport(buf.Bytes())
	if err != nil {
		return err
	}

	return nil
}

func main() {
	flag.Parse()

	if *profile < 1 || *profile > 2 {
		log.Fatalf("Invalid value for -profile: %v", *profile)
	}

	device, err := hid.Open(HoltekVendorId, AnkerMouseDeviceId, "")
	if err != nil {
		log.Fatal(err)
	}

	profileId := byte(*profile) - 1

	report1 := SetProfileReport1{
		ReportId:  0x02,
		Constant1: SetProfile1Constant1,
		Profile:   profileId,
	}

	err = WriteFeatureReport(device, report1)
	if err != nil {
		log.Fatal(err)
	}

	report2 := SetProfileReport2{
		ReportId:  0x02,
		Constant1: 0x0101,
		Profile:   profileId,
	}

	err = WriteFeatureReport(device, report2)
	if err != nil {
		log.Fatal(err)
	}
}
