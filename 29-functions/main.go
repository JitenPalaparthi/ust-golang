package main

import (
	"fmt"
	"math/rand"
)

func main() {

	var arr1 [10]int

	for i := 0; i < len(arr1); i++ {
		arr1[i] = rand.Intn(100)
	}
	fmt.Println("arr1:", arr1)
	sum1 := sumOfArr1(arr1)

	println("sum of arr1:", sum1)

	arr2 := [...]int{10, 20, 30, 40, 50} // [5]int

	// sum2 := sumOfArr1(arr2) // does not accept bcz it wants [10]int but arr2 is [5]int
	// println("sum of arr2:", sum2)
	sum2 := sumOfArr2(arr2)
	println("sum of arr2:", sum2)

	// you can convert an array to a slice

	// arr1[:] --> this is a slice not an array

	sum3, sum4 := sumOfArr(arr1[:]), sumOfArr(arr2[:])

	println("sum of arr1:", sum3)
	println("sum of arr2:", sum4)

	slice1 := arr1[:] // this is a shollow copy of array

	/*
		slice1					arr1
		-----					------
		ptr: 0x101010ff			ptr: 0x101010ff
		len: 10					len: 10
		cap: 10
	*/

	fmt.Println("arr1:", arr1)
	fmt.Println("slice1:", slice1)
	slice1[0] = 1111111111
	fmt.Println("arr1:", arr1)
	fmt.Println("slice1:", slice1)
	arr1[9] = 9999999999
	fmt.Println("arr1:", arr1)
	fmt.Println("slice1:", slice1)

	slice1 = append(slice1, 101010)
	/*
		slice1					arr1
		-----					------
		ptr: 0x202010af			ptr: 0x101010ff
		len: 11					len: 10
		cap: 20
	*/
	slice1[1] = 222222222
	fmt.Println("arr1:", arr1)
	fmt.Println("slice1:", slice1)

}

func sumOfArr1(arr [10]int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}

func sumOfArr2(arr [5]int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}

func sumOfArr(arr []int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}

// in general , dont write functions with array as a parameter
// rather write function with slice
