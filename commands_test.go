package rockblock

import (
	"fmt"
	"testing"
)

func TestQueue(t *testing.T) {

	dev := &Device{
		nil,
		"addr",
		true, false,
		nil,
		nil,
	}

	dev.initialize()

	fmt.Println(dev.writeCommand("AT+agadgad"))
	fmt.Println(dev.writeCommand("AT+fadf"))
	fmt.Println(dev.writeCommand("AT+fadf"))
}
