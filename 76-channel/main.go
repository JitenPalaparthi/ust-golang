package main

import (
	"fmt"
	"time"
)

func main() {
	// ch := make(chan int) // 8 bytes of data
	// go func() {
	// 	ch <- 100 // sender sending data <- towards the channel
	// 	fmt.Println("Hello I have sent the value")
	// }()
	// time.Sleep(time.Second * 5)
	// v := <-ch // receiver, it is from the channel
	// fmt.Println(v)

	ch := make(chan int)
	go GenerateSQ(10, ch)
	//Receiver1(10, ch)
	Receiver2(ch)
}

/*

You have a water tube that can hold 1lt of water

sending water thru the pine .. but the rereiver side (may be into a container) , the tap is closed.

*/

func GenerateSQ(num int, ch chan<- int) {
	if ch != nil {
		for i := 1; i <= num; i++ {
			time.Sleep(time.Millisecond * 100)
			ch <- i * i // this is only a sender		}
		}
		close(ch) // close sends a signal. by closing a channel the channel would not be nil. Close nothing to do with nil
	}
}

func Receiver1(num int, ch <-chan int) {
	for i := 1; i <= num; i++ {
		v := <-ch
		fmt.Println(v)
	}
}

func Receiver2(ch <-chan int) {
	for {
		v, ok := <-ch // when ever a value is received from a channel, it also sends another value that tells whether the channel is open or closed
		if ok {
			fmt.Println(v)
		} else {
			fmt.Println("Done receiving values from the channel")
			//runtime.Goexit()
			//break
			return
		}
	}
}

// ping
// pong

// ping function has to send a message to pong ("ping")
// upon receiving a message from ping, pong has to reply as pong
