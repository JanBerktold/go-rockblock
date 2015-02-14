package rockblock

import (
	"time"
)

func (dev *Device) SendMessage(msg []byte) {
}

// ReadMessage is responsible for recieving a message from the Iridium network
// and returning it to the package user. It blocks for an indefinite amount of
// time while waiting for a new message. Please note that this method does NOT
// activly check for incoming messages, but it relies on other sources to initiate
// a session. If ring alerts are enabled (default), then the user is not required to
// perform any additional requests, however if not, then he is required to call
// dev.CheckMessages() in order to refresh the message cache.
func (dev *Device) ReadMessage() Message {
	return Message{}
}

// ReadMessageWithTimeout is similar to ReadMessage, however the package user is
// able to set a maximum duration for which the method will block before returning
// without a result.
func (dev *Device) ReadMessageWithTimeout(dur time.Duration) Message {
	return Message{}
}

// CheckMessages initiates a session with the Iridium network and pulls any available
// messages down to the device before returning them to any listenting
// ReadMessage/ReadMessageWithTimeout calls.
func (dev *Device) CheckMessages() {
	dev.initiateSession()
}

func (dev *Device) checkMessage() {

}

func (dev *Device) initiateSession() {
	dev.execCommand(func() []string {
		dev.write("AT+SBDRB")
		dev.readUntil(regSesResult)

		return nil
	})
}
