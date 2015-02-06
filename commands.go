package rockblock

import (
	"regexp"
	"time"
)

type command struct {
	msg             string
	keepReg, endReg regexp.Regexp
	result          chan handleResult
}

type handleResult struct {
	msg string
	err error
}

func handleCommand(dev *Device, com *command) {
	// Simulate some work
	time.Sleep(2 * time.Second)

	if !dev.queueCommands && dev.commandCurrent != com {
		com.result <- handleResult{"", ErrCancelledTask}
	} else {
		com.result <- handleResult{com.msg + " asfasfas", nil}
	}

	if !dev.queueCommands || dev.commandQueue.Empty() {
		dev.commandWriting = false
	} else {
		handleCommand(dev, dev.commandQueue.Dequeue().(*command))
	}
}

// Function takes the command end either executes it directly or enqueues it
// It is blocking until the command is finished or stopped
func (dev *Device) writeCommand(msg string, keepReg, endReg regexp.Regexp) (string, error) {
	com := &command{
		msg,
		keepReg,
		endReg,
		make(chan handleResult),
	}

	// Lock is making sure to limit the command handling goroutines to one
	dev.commandLock.Lock()
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
	dev.commandLock.Unlock()

	result := <-com.result

	return result.msg, result.err
}
