package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	var empty struct{}
	empty = struct{}{} // no size which is zero 0 bytes

	fmt.Println("size of:", unsafe.Sizeof(empty), "type of:", reflect.TypeOf(empty))

	var r1 Rectangle
	fmt.Println("size of:", unsafe.Sizeof(r1), "type of:", reflect.TypeOf(r1))

	r1.L = 123.123
	r1.B = 123.334

	// zero values
	r2 := Rectangle{}

	fmt.Println("values of r2:", r2)

	r3 := Rectangle{L: 123.123, B: 123.23} // declare a rect variable and also assigned values to the fields
	fmt.Println("values of r3:", r3)

	r4 := Rectangle{B: 123.123, L: 123.23} // declare a rect variable and also assigned values to the fields
	fmt.Println("values of r4:", r4)

	r5 := Rectangle{123.123, 123.23} // declare a rect variable and also assigned values to the fields
	fmt.Println("values of r5:", r5)

	r6 := &Rectangle{123.123, 123.23} // r6 is pointer of rect
	//var r7 *Rectangle = &r5

	// *&r6.B = 123.123
	// *&r6.L = 123
	r6.L = 1232.123
	r6.B = 123.123
	fmt.Println("values of r6:", r6)

	r7 := new(Rectangle) // pointer using built in new function
	r7.B = 1231.123
	r7.L = 123.123

}

type Rectangle struct {
	L float32 // 0
	B float32 // 0
}
