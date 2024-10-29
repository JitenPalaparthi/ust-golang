package main

import (
	"fmt"

	"math/rand"
)

func main() {

	var arr1 [5]int  // what is the type of arr1?
	var arr2 [5]bool //

	arr3 := [3]int{10, 20, 30}         // short hand declaration
	arr4 := [...]int{12, 23, 34, 3, 4} // shorthand declaration

	//var arr5 [5]int = arr2 // cann't assign arr2 to arr5 since both types are different

	var arr5 [5]int = arr1 // can assign arr1 to arr5 since both types are same

	fmt.Println("arr1", arr1)
	fmt.Println("arr2:", arr2)
	fmt.Println("arr3:", arr3)
	fmt.Println("arr4:", arr4)
	fmt.Println("arr5:", arr5)

	for i := 0; i < len(arr1); i++ {
		arr1[i] = rand.Intn(100)
	}
	fmt.Println(arr1)

	arr2[0] = true
	arr2[1] = false
	arr2[2] = true
	arr2[3] = true
	fmt.Println(arr2)

	sum := 0

	for i := 0; i < len(arr1); i++ {
		sum += arr1[i]
	}
	println("sum of arr1:", sum)
}
