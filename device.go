// This package provides a golang interface for communicating with a rockblock device over a serial connection.
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

func connect(addr string) (*Device, error) {
	conf := &serial.Config{Name: addr, Baud: 19200}
	dev := &Device{
		nil,
		addr,
		sync.Mutex{},
	}

	if s, err := serial.OpenPort(conf); err == nil {
		dev.serial = s
		return dev, nil
	} else {
		return nil, err
	}
}

// Connect attempts to create a connection with the serial port whom the given address belongs to.
// In case of a succeeded opening, a Device object representing the connection will be returned.
// This method does not attempt to start communication with the device and therefore will suceed on
// any available serial port.
func Connect(addr string) (*Device, error) {
	return connect(addr)
}

// MustConnect functions just like Connect, however it assumes that the connection will suceed and therefore does not return an error on failure, but instead panics.
// Should be used instead of the Connect method with an ignored error parameter in order to prevent dropped errors.
func MustConnect(addr string) *Device {
	dev, err := connect(addr)
	if err != nil {
		panic(fmt.Sprintf("MustConnect (%q) failed with error %q", addr, err))
	}
	return dev
}
