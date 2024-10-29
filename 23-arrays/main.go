package main

import (
	"fmt"
	"reflect"
)

func main() {

	arr1 := [...]int{12, 23, 34, 3, 4} // shorthand declaration

	sum := 0
	for _, v := range arr1 {
		sum += v
	}

	println("sum of arr1:", sum)

	arr2d := [2][2]string{{"a", "b"}, {"c", "d"}}
	arr3d := [2][2][2]int{{{1, 2}, {3, 4}}, {{5, 6}, {7, 8}}}

	println("arr2d is a 2 dimentional array")
	for _, v1 := range arr2d {
		for _, v2 := range v1 {
			println(v2)
		}
	}
	println("arr3d is a 3 dimentional array")

	for _, v1 := range arr3d {
		for _, v2 := range v1 {
			for _, v3 := range v2 {
				println(v3)
			}
		}
	}

	fmt.Println("type of arr2d:", reflect.TypeOf(arr2d))

	fmt.Println("type of arr3d:", reflect.TypeOf(arr3d))

}

// array is not a reference type
// the length of an array should be aware to the compiler.
// array cannot be grown or shrunk at run time
