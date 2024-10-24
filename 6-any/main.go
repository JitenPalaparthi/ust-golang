package main

import (
	"fmt"
	"reflect"
)

func main() {
	var any1 any // empty interface
	fmt.Println("any1 value:", any1, "type:", reflect.TypeOf(any1))

	any1 = 1000
	fmt.Println("any1 value:", any1, "type:", reflect.TypeOf(any1))

	any1 = true
	fmt.Println("any1 value:", any1, "type:", reflect.TypeOf(any1))

	any1 = "hello world"
	fmt.Println("any1 value:", any1, "type:", reflect.TypeOf(any1))

	any1 = 123.12323
	fmt.Println("any1 value:", any1, "type:", reflect.TypeOf(any1))

	var float1 float32 = 12312.123

	any1 = float1
	fmt.Println("any1 value:", any1, "type:", reflect.TypeOf(any1))

	// var float2 float32 = float32(any1) // cannot type cast

	//var num1 int = any1.(int)           // cannot assert to int, bcz the type if any variable is float32
	// fmt.Println(float2)
	var float2 float32 = any1.(float32) // can only type assert
	fmt.Println(float2)

	var num1 int
	// any1 = 1999
	num1, ok1 := any1.(int)
	if ok1 {
		println("asserted to int", num1)
	} else {
		println("could not assert, bcz it is not int", ok1)
	}

}

// any or interface{} can hold any kind of datatype with data.
// by default any value is nil and type is also nil
// type of any variable is based on the value that is assigned
// cannot type cast any, rather a special type assertion is used
// any.(type) -> any1.(float32)
