package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	ch1, ch2 := make(chan string), make(chan string)

	go ping(ch1, ch2)

	go pong(ch2, ch1)

	runtime.Goexit() // know that this leads to runtime error

}

func ping(chPing, chPong chan string) {
	i := 1
	for {
		time.Sleep(time.Millisecond * 200)
		select {
		case chPing <- fmt.Sprint("Ping--->", i):
			i++
		case v := <-chPong:
			fmt.Println(v)
		}
	}
}

func pong(chPong, chPing chan string) {
	i := 1
	for {
		select {
		case chPong <- fmt.Sprint("Pong--->", i):
			i++
		case v := <-chPing:
			fmt.Println(v)

		}
	}
}
