package rockblock

import (
	"errors"
	"regexp"
)

var (
	ErrNoDataRecieved = errors.New("no or invalid data recieved")

	regOK         = regexp.MustCompile("^OK\r")
	regNetwork    = regexp.MustCompile("CIEV:0,[^0]")
	regTimeAnswer = regexp.MustCompile("\\+CCLK:[0-9]{2}/[0-9]{2}/[0-9]{2},[0-9]{2}:[0-9]{2}:[0-9]{2}")
)
