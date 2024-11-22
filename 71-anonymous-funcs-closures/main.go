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

	map1 := make(map[string]any)

	map1["add"] = func(i1, i2 int) int {
		return i1 + i2
	}
	map1["sub"] = sub
	map1["mul"] = mul
	map1["div"] = div
	println("\n execute map")
	a, b := 40, 30
	for k, v := range map1 {
		switch v.(type) {
		case (func(int, int) int):
			r := v.(func(int, int) int)(a, b)
			println(k, "-->", r)
		case Func:
			r1 := v.(Func)(a, b)
			println(k, "-->", r1)
			r2 := v.(Func).Rem(a, b)
			println("rem", "-->", r2)
		}
	}
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
