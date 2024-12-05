package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := new(sync.WaitGroup)
	ch := make(chan int)
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*1)
	defer cancel()
	wg.Add(1)
	go Generate(ctx, wg, ch)
	wg.Add(1)
	go func() {
		Generate(ctx, wg, ch)
	}()
	wg.Add(1)
	go func() {
	out:
		for {
			select {
			case v, ok := <-ch:
				if ok {
					fmt.Println("Received-1", v)
				} else {
					fmt.Println("ch closed")
				}
			case <-ctx.Done():
				fmt.Println("Received a cancel signal to the receiver")
				break out
			}
		}
		wg.Done()
	}()
	wg.Wait()
}

func Generate(ctx context.Context, wg *sync.WaitGroup, ch chan int) {
	i := 0
out:
	for {
		i++
		select {
		case ch <- i * i:
			time.Sleep(time.Millisecond * 200)
		case <-ctx.Done():
			println("received a signal to cancel a context.Context is getting cancelled")
			//	close(ch)
			break out
		}

	}

	wg.Done()
}
