package main

import "fmt"

func main() {

	r1 := Calc(100, 200, func(i1, i2 int) int {
		return i1 + i2
	})

	fmt.Println("sum r1:", r1)

	sub := func(i1, i2 int) int {
		return i1 - i2
	}
	r2 := Calc(123, 34, sub)
	fmt.Println("sub r2:", r2)

	r3 := Calc(123, 13, mul)

	fmt.Println("mul r3:", r3)

	var div Func = func(i1, i2 int) int {
		return i1 / i2
	}

	//r4 := Calc(123, 42, (func(int, int) int)(div)) // can type cast
	r4 := Calc(123, 42, div) // works even without
	fmt.Println("div r4:", r4)

	r5 := div.Rem(213, 23)
	fmt.Println("rem using div r5:", r5)
}

type Func func(int, int) int

func (f Func) Rem(i1, i2 int) int {
	return i1 % i2
}

func Calc(a, b int, f func(int, int) int) int {
	//r := f(a, b)
	//return r
	return f(a, b) // this is the executor of the function f
}

func mul(i, j int) int {
	return i * j
}
