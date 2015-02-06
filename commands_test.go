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
			ret, err := dev.writeCommand(start, *RegOK, *RegOK)
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

/*func BenchmarkChannelInterface(b *testing.B) {
	cha := make(chan interface{})
	err := errors.New("agagd")

	go func() {
		for i := 0; i < b.N; i++ {
			cha <- "adgadg"
			cha <- err
		}
	}()

	for i := 0; i < b.N; i++ {
		resString := (<-cha).(string)
		resError := <-cha

		// forced to use, consindering error could be nil
		var actualErr error
		if resError != nil {
			actualErr = resError.(error)
		}

		if resString == "nil" || actualErr == nil {
			fmt.Println("AGADGD FAIL")
		}
	}

}

type chanType struct {
	msg string
	err error
}

func BenchmarkChannelStruct(b *testing.B) {
	cha := make(chan chanType)
	err := errors.New("agagd")

	go func() {
		for i := 0; i < b.N; i++ {
			cha <- chanType{"adgadg", err}
		}
	}()

	for i := 0; i < b.N; i++ {
		res := <-cha

		if res.msg == "nil" || res.err == nil {
			fmt.Println("AGADGD FAIL")
		}

	}

}*/
