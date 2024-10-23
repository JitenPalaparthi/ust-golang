package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("Hello UST Global")
	fmt.Println("Current time:", time.Now())
	greet()
	os.Exit(0)
}

// func greet() {
// 	fmt.Println("Hello, I am called here , it is a greet function")
// }
