package rockblock

import (
	"fmt"
	"time"
)

func (dev *Device) waitForNetwork() {
	dev.execCommand(func() []string {
		dev.write("AT+CIER=1,1,0,0")
		dev.readUntil(regOK)
		return nil
	})
}

func (dev *Device) GetSystemTime() (time.Time, error) {
	str := dev.execCommand(func() []string {
		dev.write("AT+CCLK?")
		return dev.readUntil(regOK)
	})
	fmt.Println(str)
	if strTime := returnFirstMatch(str, regTimeAnswer); len(strTime) > 0 {
		return time.Parse("“06/01/02,15:04:05", strTime)
	}
	return time.Now(), ErrNoDataRecieved
}
