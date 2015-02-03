package rockblock

type command struct {
	msg string
}

func handleCommand(dev *Device, com command) {

	// PROGRESS AND STUFF

	// trigger handling of next enqueued command
	if dev.queueCommands {
		if dev.commandQueue.Head() != nil {
			newCom := (dev.commandQueue.Dequeue()).(command)
			handleCommand(dev, newCom)
		}
	}

}

// Function handles the execution or
func (dev *Device) writeCommand(msg string) {
	com := command{
		msg,
	}

	if dev.queueCommands {
		if dev.writingCommands {
			dev.commandQueue.Enqueue(com)
		} else {
			handleCommand(dev, com)
		}
	} else {
		// TODO: Figure out a good way to stop a command which is awaiting a reply

	}
}
