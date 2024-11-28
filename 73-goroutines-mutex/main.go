package main

import (
	"sync"
)

var (
	counter int             = 0
	wg      *sync.WaitGroup = new(sync.WaitGroup)
	mu      *sync.Mutex     = new(sync.Mutex)
)

func main() {
	wg.Add(202)
	go func() {
		for i := 1; i <= 100; i++ {
			go incr()
		}
		wg.Done()
	}()
	go func() {
		for i := 1; i <= 100; i++ {
			go decr()
		}
		wg.Done()
	}()
	wg.Wait()
	println("Counter:", counter)
}

func incr() {
	mu.Lock()
	counter++
	mu.Unlock()
	wg.Done()
}

func decr() {
	mu.Lock()
	counter--
	mu.Unlock()
	wg.Done()
}
