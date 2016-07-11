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

package device

import (
	"bytes"
	"encoding/binary"
	"github.com/GeertJohan/go.hid"
)

const (
	HoltekVendorId     = 0x04d9
	AnkerMouseDeviceId = 0xfa50
)

type Device struct {
	hiddev *hid.Device
}

func Open() (*Device, error) {
	d, err := hid.Open(HoltekVendorId, AnkerMouseDeviceId, "")
	if err != nil {
		return nil, err
	}

	return &Device{
		hiddev: d,
	}, nil
}

func (self *Device) WriteFeatureReport(report interface{}) error {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, report)
	if err != nil {
		return err
	}

	_, err = self.hiddev.SendFeatureReport(buf.Bytes())
	if err != nil {
		return err
	}

	return nil
}
