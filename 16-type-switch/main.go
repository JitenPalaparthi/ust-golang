package main

import (
	"fmt"
	"reflect"
)

func main() {

	var any1, any2 any
	var (
		num1 uint8   = 123
		num2 float64 = 123.123
	)
	any1, any2 = num1, num2
	var sum float64

	any1, any2 = 12312.123, false

	switch any1 := any1.(type) { // no need to separately type assert the variable
	case uint8:
		sum += float64(any1) // float64(any1.(uint8))
	case uint16:
		sum += float64(any1)
	case uint32:
		sum += float64(any1)
	case uint64:
		sum += float64(any1)
	case int8:
		sum += float64(any1)
	case int16:
		sum += float64(any1)
	case int32:
		sum += float64(any1)
	case int64:
		sum += float64(any1)
	case int:
		sum += float64(any1)
	case uint:
		sum += float64(any1)
	case float32:
		sum += float64(any1)
	case float64:
		sum += float64(any1)
	//default:
	//sum += 0
	default:
		println("invalid operation since the variable is not a number:", reflect.TypeOf(any1))
	}

	switch any2 := any2.(type) { // no need to separately type assert the variable
	case uint8:
		sum += float64(any2) // float64(any1.(uint8))
	case uint16:
		sum += float64(any2)
	case uint32:
		sum += float64(any2)
	case uint64:
		sum += float64(any2)
	case int8:
		sum += float64(any2)
	case int16:
		sum += float64(any2)
	case int32:
		sum += float64(any2)
	case int64:
		sum += float64(any2)
	case int:
		sum += float64(any2)
	case uint:
		sum += float64(any2)
	case float32:
		sum += float64(any2)
	case float64:
		sum += float64(any2)
	default:
		println("invalid operation since the variable is not a number:", reflect.TypeOf(any2))
		//sum += 0
	}

	fmt.Println("Sum of any1,any2:", sum)

}
