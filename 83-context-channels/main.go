package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	ch := Generate(context.TODO())
	<-Receive(context.TODO(), ch, "receiver-1--->")
	<-Receive(context.TODO(), ch, "receiver-2--->")
	time.Sleep(time.Second * 1)
}

func Generate(ctx context.Context) chan int {
	ch := make(chan int)
	go func() {
		i := 0
	out:
		for {
			select {
			case ch <- i * i:
				time.Sleep(time.Millisecond * 200)
			case <-ctx.Done():
				fmt.Println("Generator is closed as context got canceleld")
				close(ch)
				break out
			}
			i++
		}
	}()
	return ch
}

func Receive(ctx context.Context, ch <-chan int, rec string) chan struct{} {
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
			case <-ctx.Done():
				fmt.Println(rec, "is closing due to context cancellation")
				close(sig)
				break out
			}
		}
	}()
	return sig
}
