package rockblock

import (
	"errors"
	"regexp"
)

var (
	ErrCancelledTask = errors.New("AT command task has been cancelled")

	RegOK = regexp.MustCompile("^OK\r")
)

func SetQueueMode(queue bool) func(*Device) {
	return func(dev *Device) {
	}
}
