package rockblock

import (
	"regexp"
)

type Command func() []string

func (dev *Device) execCommand(com Command) []string {
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
	buf := make([]byte, 512)
	for {
		n, err := dev.serial.Read(buf)
		if err != nil {
			return result[0:i]
		} else {
			str := string(buf[0:n])
			result[i] = str
			i++
			if done.MatchString(str) {
				return result[0:i]
			}
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
