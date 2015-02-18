package rockblock

import (
	"fmt"
	"time"
)

func (dev *Device) echoOff() {
	dev.execCommand(func() []string {
		dev.write("ATE0\r")
		dev.readUntil(regOK)
		return nil
	})
}

func (dev *Device) enableRegistration() {
	dev.execCommand(func() []string {
		dev.write("AT+SBDAREG=1\r")
		dev.readUntil(regOK)
		return nil
	})
}

func (dev *Device) enableRingAlert() {
	dev.execCommand(func() []string {
		dev.write("AT+SBDMTA=1\r")
		dev.readUntil(regOK)
		return nil
	})
}

func (dev *Device) clearBuffer(ty int) {
	if ty > bufferAll {
		panic("Attempted to clear buffers with invalid type")
	}
	dev.execCommand(func() []string {
		fmt.Println(ty)
		dev.write(fmt.Sprintf("AT+SBDD%v\r", ty))
		dev.readUntil(regOK)
		return nil
	})
}

func (dev *Device) waitForNetwork() {
	dev.execCommand(func() []string {
		dev.write("AT+CIER=1,1,0,0\r")
		dev.readUntil(regOK)
		return nil
	})
}

// GetSystemTime returns the current device time, as parsed from +CCLK?.
// Request can fail in the case of bad connection or malformed replies.
func (dev *Device) GetSystemTime() (time.Time, error) {
	str := dev.execCommand(func() []string {
		dev.write("AT+CCLK?\r")
		return dev.readUntil(regOK)
	})
	fmt.Println(str)
	if strTime := returnFirstMatch(str, regTimeAnswer)[1:]; len(strTime) > 0 {
		return time.Parse("06/01/02,15:04:05", strTime[5:])
	}
	return time.Now(), ErrNoDataRecieved
}
