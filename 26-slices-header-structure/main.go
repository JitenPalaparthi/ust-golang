package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var slice1 []int // a slice is defined but not instantiated, so it is a nil slice

	/*
		slice header
		-----------
		ptr = nil
		length=0
		capacity=0
	*/

	// slice1 = make([]int, 0)
	// if slice1 == nil {
	// 	println("yes")
	// }

	//var arr1 [2]int
	if slice1 == nil {
		slice1 = make([]int, 5) // while instantiating, if you dont give the capacity, the length become the capacity
	}

	// // assign values to the slice using index
	slice1[0], slice1[1], slice1[2], slice1[3], slice1[4] = 12, 13, 45, 65, 1
	//slice1[5] = 123 // index out of range

	fmt.Println("slice1:", slice1, "length:", len(slice1), "capacity:", cap(slice1), "address:", &slice1[0])
	fmt.Printf("address of slice1:%p sizeOf slice1:%d type of slice1:%T\n", &slice1, unsafe.Sizeof(slice1), slice1)
	for i := 0; i < len(slice1); i++ {
		println(slice1[i])
	}

	// for _, v := range slice1 {
	// 	println(v)
	// }

}

/*
slice header
-----------
ptr     --> 8 bytes
length	--> 8 bytes
capacity -> 8 bytes
*/
