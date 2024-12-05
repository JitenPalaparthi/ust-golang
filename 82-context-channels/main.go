package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	parent := context.Background()
	ctx, cancel := signal.NotifyContext(parent, os.Interrupt, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGINT)
	defer cancel()
	ch := Generate(ctx)
	<-Receive(ctx, ch, "receiver-1--->")
	<-Receive(ctx, ch, "receiver-2--->")
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
				break out
			}
			i++
		}
	}()
	return ch
}

func Receive(ctx context.Context, ch <-chan int, rec string) chan struct{} {
	sig := make(chan struct{})
out:
	for {
		select {
		case v, ok := <-ch:
			if ok {
				fmt.Println(rec, v)
			} else {
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
	return sig
}
