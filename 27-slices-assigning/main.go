package main

import "fmt"

func main() {
	var slice1 []int

	if slice1 == nil {
		slice1 = make([]int, 5)
	}

	slice1[0], slice1[1], slice1[2], slice1[3], slice1[4] = 12, 13, 45, 65, 1
	fmt.Println("slice1:", slice1, "length:", len(slice1), "capacity:", cap(slice1), "address:", &slice1[0])

	slice2 := slice1 // shallow copy, refernece copy
	// header of slice1 is copied to header of slice2

	slice2[0] = 1212
	fmt.Println("slice1:", slice1, "length:", len(slice1), "capacity:", cap(slice1), "address:", &slice1[0])
	fmt.Println("slice2:", slice2, "length:", len(slice2), "capacity:", cap(slice2), "address:", &slice2[0])
	slice1 = append(slice1, 500) // slice1 and slice2 headers gets diverged
	//slice2 = slice1              // if you reassign no problem
	slice2[1] = 1300
	fmt.Println("slice1:", slice1, "length:", len(slice1), "capacity:", cap(slice1), "address:", &slice1[0])
	fmt.Println("slice2:", slice2, "length:", len(slice2), "capacity:", cap(slice2), "address:", &slice2[0])

	slice3 := make([]int, 3)

	copy(slice3, slice1) // deep copy
	fmt.Println("slice3:", slice3, "length:", len(slice3), "capacity:", cap(slice3), "address:", &slice3[0])
	slice3[0] = 10000000
	fmt.Println("slice1:", slice1, "length:", len(slice1), "capacity:", cap(slice1), "address:", &slice1[0])
	fmt.Println("slice2:", slice2, "length:", len(slice2), "capacity:", cap(slice2), "address:", &slice2[0])
	fmt.Println("slice3:", slice3, "length:", len(slice3), "capacity:", cap(slice3), "address:", &slice3[0])

	// how copy function works
	slice4 := make([]int, 3)
	l := min(len(slice1), len(slice4))

	for i := 0; i < l; i++ {
		slice4[i] = slice1[i]
	}

	fmt.Println("slice4:", slice4, "length:", len(slice4), "capacity:", cap(slice4), "address:", &slice4[0])

	clear(slice4) // clears the slice to zero value

	fmt.Println("slice4:", slice4, "length:", len(slice4), "capacity:", cap(slice4), "address:", &slice4[0])

}
