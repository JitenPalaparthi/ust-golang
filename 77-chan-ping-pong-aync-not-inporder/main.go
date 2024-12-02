package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := new(sync.WaitGroup)
	ch1, ch2 := make(chan string), make(chan string)
	wg.Add(1)
	go ping(wg, ch1, ch2)
	wg.Add(1)
	go pong(wg, ch2, ch1)
	wg.Wait()
}

func ping(wg *sync.WaitGroup, chPing, chPong chan string) {
	i := 1
	for {
		time.Sleep(time.Millisecond * 200)
		wg.Add(1)
		go func() {
			chPing <- fmt.Sprint("ping-->", i)
			i++
			wg.Done()
		}()
		wg.Add(1)
		go func() {
			fmt.Println(<-chPong)
			wg.Done()
		}()
	}
	wg.Done()
}

func pong(wg *sync.WaitGroup, chPong, chPing chan string) {
	i := 1
	for {
		wg.Add(1)
		go func() {
			fmt.Println(<-chPing)
			wg.Done()
		}()
		wg.Add(1)
		go func() {
			chPing <- fmt.Sprint("pong-->", i)
			i++
			wg.Done()
		}()
	}
	wg.Done()
}
