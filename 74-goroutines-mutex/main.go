package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	C  int
	Mu *sync.Mutex
}

func New() *Counter {
	return &Counter{C: 0, Mu: new(sync.Mutex)}
}

func (c *Counter) Incr() {
	c.Mu.Lock()
	c.C++
	c.Mu.Unlock()
}

func (c *Counter) Decr() {
	c.Mu.Lock()
	c.C--
	c.Mu.Unlock()
}
func (c *Counter) PrintCount() {
	fmt.Println(c.C)
}

func main() {

	c := New()
	wg := new(sync.WaitGroup)

	wg.Add(1)
	go func() {
		for i := 1; i <= 100; i++ {
			wg.Add(1)
			go func() {
				c.Incr()
				wg.Done()
			}()
		}
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		for i := 1; i <= 100; i++ {
			wg.Add(1)
			go func() {
				c.Decr()
				wg.Done()
			}()
		}
		wg.Done()
	}()
	wg.Wait()
	c.PrintCount()

}
