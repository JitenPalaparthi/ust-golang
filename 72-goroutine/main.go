package main

import (
	"fmt"
	"sync"
)

var count int = 0

func main() {
	wg := new(sync.WaitGroup)
	defer println("end of main")
	for i := 1; i <= 50; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			greet1(i)
			wg.Done()
		}(wg)
		wg.Add(1)
		go greet2(wg, i)
	}
	wg.Add(1)
	var s int
	go func() {
		s = add(100, 200)
		fmt.Println("Sum s:", s)
		wg.Done()
	}()
	wg.Wait()
	//time.Sleep(time.Millisecond * 1)
	fmt.Println(count)
	//runtime.Goexit() // main does not return since it gives exit non zero, exit non zero is an error
	// This waits until the counter becomes zero
}

func greet1(i int) {
	count++
	println("Hello World-->", fmt.Sprint(i))
	// This has completed its execution
}

func greet2(wg *sync.WaitGroup, i int) {
	count++
	println("Hello World-->", fmt.Sprint(i))
	defer println(count)
	wg.Done()
	// This has completed its execution
}

func add(i, j int) int {
	return i + j
}

// func evenNumbers(i, j int) {

// }

// WaitGroup
// WaitGroup has a counter
// For every Goroutine before start it has to be added to the counter
// Everytime A Goroutine completes its execution wg.Done , the counter decreases
// have to wait the main, not to return until the Counter becomes zero
//
