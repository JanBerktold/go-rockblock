package rockblock

import (
	"fmt"
	"github.com/oleiade/lane"
	"github.com/tarm/goserial"
	"io"
)

type Device struct {
	serial io.ReadWriteCloser
	addr   string

	// AT Command handling
	writingCommands bool
	queueCommands   bool
	commandQueue    *lane.Queue
	commandCurrent  *command
}

// Internal function to load required dependencies based upon configuration
func (dev *Device) initialize() {

	if dev.queueCommands {
		dev.commandQueue = lane.NewQueue()
	}

}

func connect(addr string, options []func(*Device)) (*Device, error) {
	conf := &serial.Config{Name: addr, Baud: 19200}
	dev := &Device{
		nil,
		addr,
		false,
		true,
		nil,
		nil,
	}

	// apply user options
	for _, fun := range options {
		fun(dev)
	}

	dev.initialize()

	if s, err := serial.OpenPort(conf); err == nil {
		dev.serial = s
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
