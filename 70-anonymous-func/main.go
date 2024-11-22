package main

import "fmt"

func main() {
	var f1 func(int, int) int // variables can also be funcs
	// f1 can be nil? what does it mean?

	if f1 == nil {
		fmt.Printf("yes f1 is nil, address %p\n", f1)
	}

	func() { // input and output types
		println("Hello UST Global")
	}() // the executor

	c := func(a, b int) int {
		return a + b
	}(10, 20)
	println("c:", c)

	f1 = func(a, b int) int {
		return a - b
	}
	if f1 == nil {
		println("func f is nil")
	} else {
		fmt.Printf("yes f1 is not nil, address %p\n", f1)
	}

	d := f1(123, 12)
	fmt.Println("d:", d)

	//{
	num := 1

	func() {
		for i := 1; i <= 100; i++ {
			num++
		}
	}()
	println("num:", num)
	//}

	var f2 Func1 = func(i1, i2 int) int {
		return i1 * i2
	}
	r1 := f2(12, 123)
	f2.Hey()
	fmt.Println("r1:", r1)
}

type Func1 func(int, int) int

func (f Func1) Hey() {
	fmt.Println("Hello, just checking how a method can be attached to a func")
}

// 1. anonymous functions , funcs without names
// 2. funcs can be creaeted as variables
// 3. closures can be created using anonymous functions
// 4. expressions can evaluated from scopes using anonymous funcs.
//		Since scopes cannot evaluate expressiomns in go unlike in rust
