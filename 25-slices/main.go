package main

import "fmt"

func main() {
	var slice1 []int // a slice is defined but not instantiated, so it is a nil slice

	if slice1 == nil {
		slice1 = make([]int, 5) // while instantiating, if you dont give the capacity, the length become the capacity
	}

	// assign values to the slice using index
	slice1[0], slice1[1], slice1[2], slice1[3], slice1[4] = 12, 13, 45, 65, 1
	//slice1[5] = 123 // index out of range

	fmt.Println("slice1:", slice1, "length:", len(slice1), "capacity:", cap(slice1), "address:", &slice1[0])

	// to grow a slice, there is a append built in function

	slice1 = append(slice1, 100) // the initial capasity is 5 but now appending the slice to the 6th element
	// would increases the capacity hence it reallocates new memory and gives more capacity
	// hence the new reference to be given

	fmt.Println("slice1:", slice1, "length:", len(slice1), "capacity:", cap(slice1), "address:", &slice1[0])
	slice1 = append(slice1, 200)
	fmt.Println("slice1:", slice1, "length:", len(slice1), "capacity:", cap(slice1), "address:", &slice1[0])
	slice1 = append(slice1, 300, 400, 500) // total length is 10 and capacity 10
	fmt.Println("slice1:", slice1, "length:", len(slice1), "capacity:", cap(slice1), "address:", &slice1[0])
	slice1 = append(slice1, 600)
	// 1. it reallocates to a new memory region
	// 2. all existing elements in the slice to be copied to the new memore
	// 3. the previous memory to be deallocated(is no references are attached to it), this is done by GC
	fmt.Println("slice1:", slice1, "length:", len(slice1), "capacity:", cap(slice1), "address:", &slice1[0])

	sum := 0

	for _, v := range slice1 {
		sum += v
	}

	fmt.Println("Sum of slice1:", sum)
}
