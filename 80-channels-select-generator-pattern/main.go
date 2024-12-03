package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := Generate(time.Second*2, time.Millisecond*500) // generator pattern
	//signal := <-Receiver(ch)
	//<-Receiver(ch1) // future pattern
	//fmt.Println("Received a signal, so exiting from main", signal)
	<-Receiver(ch1)

	ch2 := Generate(time.Second*1, time.Millisecond*200)
	<-Receiver(ch2)
	fmt.Println("Received a signal, so exiting from main")
}

func Generate(cancelDuration time.Duration, sleepDuration time.Duration) chan int {
	i := 1
	timeOutch := timeOutAfter(cancelDuration)
	ch := make(chan int)
	go func() {
	out:
		for {
			select {
			case ch <- i * i:
				time.Sleep(sleepDuration)
				i++
			case _, ok := <-timeOutch:
				if !ok {
					close(ch)
					break out
				}
			}
		}
	}()
	return ch
}

func Receiver(ch <-chan int) chan struct{} {
	signal := make(chan struct{})
	go func() {
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
	}()
	return signal
}

func timeOutAfter(d time.Duration) chan bool {
	timeOutCh := make(chan bool)
	go func() {
		time.Sleep(d)
		timeOutCh <- true
		close(timeOutCh)
	}()
	return timeOutCh
}
