package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	ch := make(chan int)
	closed := make(chan struct{})
	go Generate(ch, closed)
	go Generate(ch, closed)
	<-Receive(ch, "receiver-1--->")
	time.Sleep(time.Second * 1)
}

func Generate(ch chan<- int, closed chan struct{}) {
	//timeOutch := time.After(d)
	go func() {
		i := 1
		for {
			if i > 10 {
				close(closed)
				runtime.Goexit()
			}
			select {
			case ch <- i * i:
			case <-closed:
				close(ch)
				runtime.Goexit()
			}
			i++
		}
		//close(closed)
	}()

}

func Receive(ch <-chan int, rec string) chan struct{} {
	sig := make(chan struct{})
	go func() {
	out:
		for {
			select {
			case v, ok := <-ch:
				if ok {
					fmt.Println(rec, v)
				} else {
					fmt.Println("Closing this channel on --", rec)
					sig <- struct{}{}
					close(sig)
					break out
				}
			}
		}
	}()
	return sig
}
