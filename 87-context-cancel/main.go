package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	wg := new(sync.WaitGroup)
	ch := make(chan int, 1)
	ctx := context.Background()
	// ctx, cancel := context.WithTimeout(ctx, time.Second*1)
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM, syscall.SIGKILL)
	defer cancel()
	wg.Add(1)
	go func() {
		ctxp, cancel := context.WithCancel(ctx)
		defer cancel()
		i := 0
	out:
		for {
			i++
			select {
			case ch <- i * i:
				time.Sleep(time.Millisecond * 200)
			case v, ok := <-ch:
				if ok {
					fmt.Println(v)
				}
			case <-ctxp.Done():
				println("received a signal to cancel a context.Context is getting cancelled in generate")
				//	close(ch)
				break out
			}

		}
		wg.Done()
	}()
	wg.Wait()
}

func Generate(ctx context.Context, wg *sync.WaitGroup, ch chan int) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()
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
