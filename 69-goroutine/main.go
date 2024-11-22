package main

import (
	"fmt"
	"runtime"
)

func main() {
	for i := 1; i <= 100; i++ {
		go Greet(i)
	}
	println("Hello UST Folks")
	//time.Sleep(time.Microsecond * 2) //blocking the thread
	runtime.Goexit() // exits the goroutine
}

func Greet(i int) {
	fmt.Println("hello World-->", i)
	//runtime.Goexit()
}

// 1. Main is also a goroutine
// 2. By default no goroutine waits for other goroutines to finish their execution
// 3. Cannot guarantee the order of execution
// 4. to create a goroutine just use a go keyword infront of it
