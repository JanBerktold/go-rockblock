package rockblock

func SetQueueMode(queue bool) func(*Device) {
	return func(dev *Device) {
		dev.QueueCommands = queue
	}
}
