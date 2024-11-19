package main

import "fmt"

func main() {

	defer func() {
		fmt.Println("Hello World")
	}()

	defer func(msg string) {
		fmt.Println(msg)
	}("Hello UST Global minds")

	c := func(a, b int) int {
		return a + b
	}(10, 20)

	fmt.Println("c:", c)

	num := 6

	defer func() {
		sq := func() int {
			return num * num
		}()

		fmt.Println("squre of num:", sq)
	}()

	f := func(a, b int) int {
		return a - b
	}

	c1 := f(100, 114)

	fmt.Println("c1:", c1)

}
