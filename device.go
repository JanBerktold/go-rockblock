package rockblock

import (
	"fmt"
	"github.com/tarm/goserial"
	"io"
)

type Device struct {
	serial io.ReadWriteCloser
	addr   string
}

func Connect(addr string, options ...func(*Device)) (*Device, error) {
	c := &serial.Config{Name: addr, Baud: 19200}
	s, err := serial.OpenPort(c)

	if err == nil {
		dev := &Device{s, addr}

		// apply user options
		// perhaps move this prior to initalization of the serial port?
		for _, fun := range options {
			fun(&dev)
		}

		return dev, nil
	} else {
		return nil, err
	}
}

func MustConnect(addr string, options ...func(*Device)) *Device {
	dev, err := Connect(addr, options)
	if err != nil {
		panic(fmt.Sprintf("MustConnect (%q) failed with error %q", addr, err))
	}
	return dev
}
