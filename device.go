package rockblock

import (
	"fmt"
	"github.com/tarm/goserial"
	"io"
	"sync"
)

type Device struct {
	serial io.ReadWriteCloser
	addr   string

	commandLock sync.Mutex
}

func connect(addr string, options []func(*Device)) (*Device, error) {
	conf := &serial.Config{Name: addr, Baud: 19200}
	dev := &Device{
		nil,
		addr,
		sync.Mutex{},
	}

	// apply user options
	for _, fun := range options {
		fun(dev)
	}

	if s, err := serial.OpenPort(conf); err == nil {
		dev.serial =
			s
		return dev, nil
	} else {
		return nil, err
	}
}

func Connect(addr string, options ...func(*Device)) (*Device, error) {
	return connect(addr, options)
}

func MustConnect(addr string, options ...func(*Device)) *Device {
	dev, err := connect(addr, options)
	if err != nil {
		panic(fmt.Sprintf("MustConnect (%q) failed with error %q", addr, err))
	}
	return dev
}
