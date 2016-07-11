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
	"fmt"
	"github.com/GeertJohan/go.hid"
	colorful "github.com/lucasb-eyer/go-colorful"
)

var (
	configuration = [][]byte{
		{2, 6, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{2, 2, 16, 0, 8, 0, 250, 250, 226, 48, 151, 47, 207, 187, 228, 64},
		{2, 3, 64, 0, 1, 0, 250, 250, 0, 0, 0, 0, 0, 0, 0, 0},
		{2, 2, 69, 0, 1, 0, 250, 250, 2 /* hz? */, 0, 0, 0, 0, 0, 0, 0},
		{2, 3, 72, 0, 32, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0 /* buttons profile 1 */},
		{3, 2, 209, 0, 21, 0, 250, 250, 129, 1, 1, 6, 1, 0, 1, 1, 1, 6, 2, 0, 129, 1, 1, 6, 1, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0 /* light profile 1 */},
		{0 /* DPI profile 1 */},
		{0 /* buttons profile 2 */},
		{3, 2, 209, 9, 21, 0, 250, 250, 129, 1, 1, 6, 1, 0, 1, 1, 1, 6, 2, 0, 129, 1, 1, 6, 1, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0 /* light profile 2 */},
		{0 /* DPI profile 2 */},
		{2, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, /* Commit? */
	}
)

const (
	ButtonsProfile1Idx = 5
	LightProfile1Idx   = 7
	DPIProfile1Idx     = 8
	ButtonsProfile2Idx = 9
	LightProfile2Idx   = 11
	DPIProfile2Idx     = 12
)

const (
	EventDisabled    = 0x0e
	EventLeftClick   = 0x01
	EventRightClick  = 0x02
	EventMiddleCLick = 0x03
	EventForward     = 0x05
	EventSingleKey   = 0x10
)

type ButtonEntry struct {
	EventId      byte   // Event* constants above
	ExtendedInfo byte   // Still-unclear
	KeyId        uint16 // USB Scancodes if EventSingleKey

}

type ButtonsProfile struct {
	ReportId   byte    // 0x04
	InternalId byte    // 0x02
	Constant1  byte    // 0x90
	ProfileId  byte    // Profile1=0x00 Profile2=0x09
	Constant2  [5]byte // 0x41 0x00 0xFA 0xFA 0x10
	Buttons    [9]ButtonEntry
	Constant3  byte      // 0x0d
	Unknown    [978]byte // All zeroes
}

func NewButtonsProfile(profile int) *ButtonsProfile {
	var profileId byte

	switch profile {
	case 1:
		profileId = 0x00
	case 2:
		profileId = 0x09
	}

	return &ButtonsProfile{
		ReportId:   0x04,
		InternalId: 0x02,
		Constant1:  0x90,
		ProfileId:  profileId,
		Constant2:  [5]byte{0x41, 0x00, 0xFA, 0xFA, 0x10},
		Buttons: [9]ButtonEntry{
			// Default mapping from the Anker app (In parenthesis, the original button number)
			{0x01, 0x00, 0x0000}, // (1) Left Click
			{0x02, 0x00, 0x0000}, // (2) Right Click
			{0x03, 0x00, 0x0000}, // (3) Middle Click
			{0x05, 0x00, 0x0000}, // (4) Forward
			{0x10, 0x00, 0x00E2}, // (6) L ALT
			{0x04, 0x00, 0x0000}, // (5?) Back?
			{0x11, 0x00, 0x0015}, // (7?) Macro Play?
			{0x08, 0x00, 0x0000}, // (8?) Marco Record?
			{0x13, 0x80, 0x0000}, // (9?) DPI Switch?
		},
		Constant3: 0x0d,
	}
}

func (self *ButtonsProfile) ToBytes() ([]byte, error) {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, self)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type LightProfile struct {
	ReportId     byte   // 0x02
	InternalId   byte   // 0x02
	Constant1    byte   // 0x81
	ProfileId    byte   // Profile1=0x08 Profile2=0x11
	Constant3    byte   // 0x06
	Constant4    byte   // 0x00
	Constant5    uint16 // 0xFAFA
	InverseRed   byte
	InverseGreen byte
	InverseBlue  byte
	Brightness   byte
	BreathSpeed  byte
	Constant6    [3]byte // 0x00 0x00 0x00
}

func NewLightProfile(profile int) *LightProfile {
	var profileId byte

	switch profile {
	case 1:
		profileId = 0x08
	case 2:
		profileId = 0x11
	}

	return &LightProfile{
		ReportId:   0x02,
		InternalId: 0x02,
		Constant1:  0x81,
		ProfileId:  profileId,
		Constant3:  0x06,
		Constant5:  0xFAFA,
	}
}

func (self *LightProfile) SetColor(c colorful.Color) {
	r, g, b := c.RGB255()
	self.InverseRed = ^r
	self.InverseGreen = ^g
	self.InverseBlue = ^b
}

func (self *LightProfile) ToBytes() ([]byte, error) {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, self)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type dpiEntry struct {
	Enabled byte // enabled=1 disabled=0
	X       byte // 50 dpi for each unit (i.e. 20 = 1000dpi)
	Y       byte
}

type DPIProfile struct {
	ReportId   byte    // 0x03
	InternalId byte    //0x02
	ProfileId  byte    // Profile1=0x00 Profile2=0x09
	Constant1  [6]byte // 0x20 0x00 0xFA 0xFA 0x04 0x01 (last one not sure is a constant)
	DPI        [4]dpiEntry
	Unknown    [42]byte // All Zeroes
}

var DPIProfileConstant1 = [6]byte{0x20, 0x00, 0xFA, 0xFA, 0x04, 0x01}

func NewDPIProfile(profile int) *DPIProfile {
	var profileId byte

	switch profile {
	case 1:
		profileId = 0x00
	case 2:
		profileId = 0x09
	}

	return &DPIProfile{
		ReportId:   0x03,
		InternalId: 0x02,
		ProfileId:  profileId,
		Constant1:  DPIProfileConstant1,
		DPI: [4]dpiEntry{
			{Enabled: 1, X: 20, Y: 20},
			{Enabled: 1, X: 40, Y: 40},
			{Enabled: 1, X: 80, Y: 80},
			{Enabled: 1, X: 164, Y: 164},
		},
	}
}

func (self *DPIProfile) SetDPI(dpi [4][2]int) {
	for i := range self.DPI {
		// If the DPI value is 0, we consider it disabled.
		if dpi[i][0] == 0 {
			self.DPI[i].Enabled = 0
		} else {
			self.DPI[i] = dpiEntry{
				Enabled: 1,
				X:       byte(dpi[i][0] / 50),
				Y:       byte(dpi[i][1] / 50),
			}
		}
	}
}

func (self *DPIProfile) ToBytes() ([]byte, error) {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, self)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type ConfigProfile struct {
	*ButtonsProfile
	*LightProfile
	*DPIProfile
}

func NewConfigProfile(profile int) *ConfigProfile {
	return &ConfigProfile{
		ButtonsProfile: NewButtonsProfile(profile),
		LightProfile:   NewLightProfile(profile),
		DPIProfile:     NewDPIProfile(profile),
	}
}

type Config struct {
	Profiles [2]*ConfigProfile
}

func NewConfig() *Config {
	return &Config{
		Profiles: [2]*ConfigProfile{
			NewConfigProfile(1),
			NewConfigProfile(2),
		},
	}
}

func (self *Config) Write(device *hid.Device) error {
	reports := configuration

	var r []byte
	var err error

	r, err = self.Profiles[0].ButtonsProfile.ToBytes()
	if err != nil {
		return err
	}
	reports[ButtonsProfile1Idx] = r

	r, err = self.Profiles[0].LightProfile.ToBytes()
	if err != nil {
		return err
	}
	reports[LightProfile1Idx] = r

	r, err = self.Profiles[0].DPIProfile.ToBytes()
	if err != nil {
		return err
	}
	reports[DPIProfile1Idx] = r

	r, err = self.Profiles[0].ButtonsProfile.ToBytes()
	if err != nil {
		return err
	}
	reports[ButtonsProfile2Idx] = r

	r, err = self.Profiles[1].LightProfile.ToBytes()
	if err != nil {
		return err
	}
	reports[LightProfile2Idx] = r

	r, err = self.Profiles[1].DPIProfile.ToBytes()
	if err != nil {
		return err
	}
	reports[DPIProfile2Idx] = r

	for i, r := range reports {
		_, err := device.SendFeatureReport(r)
		if err != nil {
			return fmt.Errorf("Error writing report %v: %v", i, err)
		}
	}

	return nil
}
