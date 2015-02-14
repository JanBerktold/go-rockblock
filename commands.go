package rockblock

import (
	"fmt"
	"regexp"
)

type command func() []string

func (dev *Device) pullMessages() {
	buf := make([]byte, 512)
	for {
		if n, err := dev.serial.Read(buf); err == nil {
			msg := string(buf[0:n])
			// check for unsolicited messages
			if regSbRing.MatchString(msg) {
				fmt.Println(msg)
				dev.initiateSession()
			} else if regAreg.MatchString(msg) {
				fmt.Println(msg)
			} else {
				dev.serialChannel <- msg
			}
		} else {
			dev.serialChannel <- ""
			return
		}
	}
}

func (dev *Device) execCommand(com command) []string {
	dev.commandLock.Lock()
	defer dev.commandLock.Unlock()
	return com()
}

func (dev *Device) write(str string) {
	dev.serial.Write([]byte(str))
}

func (dev *Device) readUntil(done *regexp.Regexp) []string {
	result := make([]string, 10)
	i := 0
	for {
		str := <-dev.serialChannel
		if len(str) > 0 {
			result[i] = str
			i++
			if done.MatchString(str) {
				return result[0:i]
			}
		} else {
			return result[0:i]
		}
	}
}

func returnFirstMatch(str []string, reg *regexp.Regexp) string {
	for _, str := range str {
		if reg.MatchString(str) {
			return str
		}
	}
	return ""
}
