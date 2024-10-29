package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {

	var any1 any
	// what is the type and value inference of any variable
	// by default if no value or type is assigned to any variable, both are nil and nil
	println(unsafe.Sizeof(any1))
	any1 = 1231.123
	println(unsafe.Sizeof(any1))
	any1 = "Hello World"
	println(unsafe.Sizeof(any1))
	any1 = true
	println(unsafe.Sizeof(any1))
	// above statements is to explain why any is also a statically typed type

	//println(any1.(type)) any1.(type) it is specially designed to work with switch statement
	var num1 uint8 = 123
	any1 = num1
	switch any1.(type) {
	case int:
		fmt.Println("Value:", any1, "Type:", reflect.TypeOf(any1))
	case float64:
		fmt.Println("Value:", any1, "Type:", reflect.TypeOf(any1))
	case bool:
		fmt.Println("Value:", any1, "Type:", reflect.TypeOf(any1))
	case string:
		fmt.Println("Value:", any1, "Type:", reflect.TypeOf(any1))
	default:
		fmt.Println("none of the types of int,float64,bool or string.any1 belongs to", reflect.TypeOf(any1), "type")
	}

}

// Go is statically type language --> Does this true for any as well?
