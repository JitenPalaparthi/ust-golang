package main

import (
	"errors"
	"fmt"
)

//var G int

func main() {
	var num1 int = 100
	var ptr1 *int = &num1
	println("direct: address of num1:", &num1, "value in num1:", num1) // dereference
	println("pointer:address of num1:", ptr1, "value in num1:", *ptr1) // dereference

	Increment(num1)

	println("num1 after the increment:", num1)

	IncrementP(&num1)
	println("num1 after the increment:", num1)

	IncrementP(ptr1)
	println("num1 after the increment:", num1)

	num1 = 300
	*&num1 = 400
	println("num1:", num1)

	// ptr1++ // pointer arthemetic is not allowed in go (directly)

	var ptr2 *int
	IncrementP(ptr2)

	if err := IncrementPE(ptr2); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Value in ptr2 address", *ptr2)
	}

}

func Increment(num int) {
	num++
}

func IncrementP(num *int) {
	if num == nil {
		return
	}
	*num++
}

func IncrementPE(num *int) error {
	if num == nil {
		return errors.New("nil pointer")
	}
	*num++
	return nil
}
