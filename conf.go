package rockblock

import (
	"errors"
)

var (
	ErrCancelledTask = errors.New("AT command task has been cancelled")
)

func SetQueueMode(queue bool) func(*Device) {
	return func(dev *Device) {
		dev.queueCommands = queue
	}
}
