package rockblock

import (
	"bytes"
	"regexp"
)

type command func() []string

func (dev *Device) pullMessages() {
	buf := make([]byte, 512)
	command := bytes.NewBufferString("")
	for {
		if n, err := dev.serial.Read(buf); err == nil {
			msg := string(buf[0:n])

			for _, str := range msg {
				if str == '\r' {
					actualMessage := command.String()

					if regSbRing.MatchString(actualMessage) {
						dev.initiateSession()
					} else if regAreg.MatchString(actualMessage) {
					} else {
						dev.serialChannel <- actualMessage
					}

					command.Reset()
				} else if str != '\n' {
					command.WriteRune(str)
				}
			}
		} else {
			dev.serialChannel <- "end_this"
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
	result := make([]string, 255)
	i := 0
	for {
		str := <-dev.serialChannel
		if str != "end_this" {
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
