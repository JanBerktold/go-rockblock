package rockblock

type command struct {
	msg string
}

func handleCommand(dev *Device, com command) {

}

// Function handles the execution or
func (dev *Device) writeCommand(msg string) {

	// Function takes the command end either executes it directly or enqueues it
	// It is blocking until the command is finished or stopped

}
