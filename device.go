// This package provides a golang interface for communicating with a rockblock device over a serial connection.
package rockblock

import (
	"fmt"
	"github.com/tarm/goserial"
	"io"
	"sync"
)

type Message struct {
}

type Device struct {
	serial io.ReadWriteCloser
	addr   string

	serialChannel chan string
	commandLock   sync.Mutex
}

func connect(addr string) (*Device, error) {
	conf := &serial.Config{Name: addr, Baud: 19200}
	dev := &Device{
		nil,
		addr,
		make(chan string),
		sync.Mutex{},
	}

	if s, err := serial.OpenPort(conf); err == nil {
		dev.serial = s

		dev.echoOff()
		dev.clearBuffer(bufferAll)
		dev.enableRegistration()
		dev.enableRingAlert()

		return dev, nil
	} else {
		return nil, err
	}
}

// Connect attempts to create a connection with the serial port whom the given address belongs to as well as
// performing original setup procedures. In case of a succeeded opening, a Device object representing the
// connection will be returned.
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
