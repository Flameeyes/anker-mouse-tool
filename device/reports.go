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

package device

type SetLightReport struct {
	reportId     byte // 0x02
	internalId   byte // 0x04
	InverseRed   byte
	InverseGreen byte
	InverseBlue  byte
	Brightness   byte
	BreathSpeed  byte
	unknown      [9]byte // All zeroes.
}

func newSetLightReport(r, g, b, brightness, breathSpeed byte) *SetLightReport {
	return &SetLightReport{
		reportId:     0x02,
		internalId:   0x04,
		InverseRed:   ^r,
		InverseGreen: ^g,
		InverseBlue:  ^b,
		Brightness:   brightness,
		BreathSpeed:  breathSpeed,
	}
}

type setProfileReport1 struct {
	reportId  byte    // 0x02
	constant1 [7]byte // 0x02, 0x40, 0x00, 0x01, 0x00, 0xFA, 0xFA
	Profile   byte
	unknown   [7]byte // All zeroes.
}

var setProfile1Constant1 = [7]byte{0x02, 0x40, 0x00, 0x01, 0x00, 0xFA, 0xFA}

type setProfileReport2 struct {
	reportId  byte   // 0x02
	constant1 uint16 // 0x0101
	Profile   byte
	unknown   [12]byte // All zeroes.
}

func newSetProfileReports(byte profileId) (*setProfileReport1, *setProfileReport2) {
	r1 := setProfileReport1{
		reportId:  0x02,
		constant1: setProfile1Constant1,
		Profile:   profileId,
	}

	r2 := setProfileReport2{
		reportId:  0x02,
		constant1: 0x0101,
		Profile:   profileId,
	}

	return &r1, &r2
}
