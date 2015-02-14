package rockblock

import (
	"time"
)

func (dev *Device) SendBinaryMessage(msg []byte) {
}

func (dev *Device) SendMessage(msg string) {
}

// ReadMessage is responsible for recieving a message from the Iridium network
// and returning it to the package user. It blocks for an indefinite amount of
// time while waiting for a new message. Please note that this method does NOT
// activly check for incoming messages, but it relies on other sources to initiate
// a session.
func (dev *Device) ReadMessage() Message {
	return Message{}
}

// ReadMessageWithTimeout is similar to ReadMessage, however the package user is
// able to set a maximum duration for which the method will block before returning
// without a result.
func (dev *Device) ReadMessageWithTimeout(dur time.Duration) Message {
	return Message{}
}

func (dev *Device) initiateSession() {

}
