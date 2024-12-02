package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	//signal := make(chan bool)
	signal := make(chan struct{})
	go Generate(ch, time.Second*1, time.Millisecond*200)
	go Receiver(ch, signal)
	<-signal
}

func Generate(ch chan<- int, cancelDuration time.Duration, sleepDuration time.Duration) {
	i := 1
	//timeOutCh := time.After(cancelDuration) // after the given duration, this function returns a channel
	//timeOutCh := make(chan bool)
	//timeOutAfter(cancelDuration, timeOutCh)
	timeOutch := timeOutAfterR(cancelDuration)
out:
	for {
		select {
		case ch <- i * i:
			time.Sleep(sleepDuration)
			i++
		// case <-timeOutCh:
		// 	close(ch)
		// 	break out
		case _, ok := <-timeOutch:
			if !ok {
				close(ch)
				break out
			}
		}
	}
}

// func Receiver(ch <-chan int, signal chan<- bool) {

// out:
// 	for {
// 		select {
// 		case v, ok := <-ch:
// 			if ok {
// 				fmt.Println(v)
// 			} else {
// 				signal <- true
// 				close(signal)
// 				break out
// 			}
// 		}
// 	}
// }

func Receiver(ch <-chan int, signal chan<- struct{}) {
out:
	for {
		select {
		case v, ok := <-ch:
			if ok {
				fmt.Println(v)
			} else {
				signal <- struct{}{}
				close(signal)
				break out
			}
		}
	}
}

func timeOutAfter(d time.Duration, chSignal chan bool) {
	time.Sleep(d)
	chSignal <- true
}

func timeOutAfterR(d time.Duration) chan bool {
	timeOutCh := make(chan bool)
	go func() {
		time.Sleep(d)
		timeOutCh <- true
		close(timeOutCh)
	}()
	return timeOutCh
}

// empty struct
