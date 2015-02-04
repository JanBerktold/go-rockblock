package rockblock

import "time"

type command struct {
	msg    string
	result chan interface{}
}

func handleCommand(dev *Device, com *command) {
	// Simulate some work
	time.Sleep(5 * time.Second)

	if !dev.queueCommands && dev.commandCurrent != com {
		com.result <- ""
		com.result <- ErrCancelledTask
	} else {
		com.result <- com.msg + " asfasfas"
		com.result <- nil
	}

	if !dev.queueCommands || dev.commandQueue.Empty() {
		dev.commandWriting = false
	} else {
		handleCommand(dev, dev.commandQueue.Dequeue().(*command))
	}
}

// Function handles the execution or
func (dev *Device) writeCommand(msg string) (string, error) {

	// Function takes the command end either executes it directly or enqueues it
	// It is blocking until the command is finished or stopped

	com := &command{
		msg,
		make(chan interface{}),
	}

	if dev.queueCommands {
		if dev.commandWriting {
			dev.commandQueue.Enqueue(com)
		} else {
			dev.commandWriting = true
			go handleCommand(dev, com)
		}
	} else {
		// !! NOT TESTED
		dev.commandCurrent = com
		go handleCommand(dev, com)
	}

	resString := (<-com.result).(string)
	resError := <-com.result

	// forced to use, consindering error could be nil
	var actualErr error
	if resError != nil {
		actualErr = resError.(error)
	}

	return resString, actualErr
}
