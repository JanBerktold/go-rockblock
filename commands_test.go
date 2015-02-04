package rockblock

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestQueue(t *testing.T) {

	dev := &Device{
		nil,
		"addr",
		sync.Mutex{},
		true, false,
		nil,
		nil,
	}

	dev.initialize()

	// ugly
	cha := make(chan bool)
	for i := 0; i < 5; i++ {
		go func() {
			start := fmt.Sprintf("AT+TESTHG%v", i)
			ret, err := dev.writeCommand(start)
			cha <- true
			if ret != start+" asfasfas" || err != nil {
				t.FailNow()
			}
		}()
	}

	// timeout
	go func() {
		time.Sleep(12 * time.Second)
		cha <- false
	}()

	// check for completed queue tasks
	for i := 0; i < 5; i++ {
		if !(<-cha) {
			t.FailNow()
		}
	}

}
