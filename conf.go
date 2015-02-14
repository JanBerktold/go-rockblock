package rockblock

import (
	"errors"
	"regexp"
)

var (
	ErrNoDataRecieved = errors.New("no or invalid data recieved")

	RegOK         = regexp.MustCompile("^OK\r")
	RegNetwork    = regexp.MustCompile("CIEV:0,[^0]")
	RegTimeAnswer = regexp.MustCompile("\\+CCLK:[0-9]{2}/[0-9]{2}/[0-9]{2},[0-9]{2}:[0-9]{2}:[0-9]{2}")
)

func SetQueueMode(queue bool) func(*Device) {
	return func(dev *Device) {
	}
}
