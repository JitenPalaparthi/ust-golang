package main

import "fmt"

func main() {
	slice1 := make([]int, 5, 10)
	slice1[0], slice1[1], slice1[2], slice1[3], slice1[4] = 1, 2, 3, 4, 5
	fmt.Println("slice1:", slice1)
	addElementsAndDoubleAll(slice1)
	fmt.Println("slice1:", slice1)
	addElementsAndDoubleAll(slice1, 6, 7, 8, 9, 10)
	fmt.Println("slice1:", slice1)
	slice1 = addElementsAndDoubleAll(slice1, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15)
	fmt.Println("slice1:", slice1)

	slice2 := slice1[2:] // from the 2nd element to the end is assigned to slice2
	fmt.Println("slice2:", slice2, "len", len(slice2), "cap", cap(slice2))
	slice3 := slice1[:10] // from the 0 to 10 but not 10
	fmt.Println("slice3:", slice3, "len", len(slice3), "cap", cap(slice3))
	slice4 := slice1[5:10] // from the 5 to 10 but not 10
	fmt.Println("slice4:", slice4, "len", len(slice4), "cap", cap(slice4))

}

func addElementsAndDoubleAll(slice []int, elems ...int) []int {
	slice = append(slice, elems...)
	for i, v := range slice {
		slice[i] = v * 2
	}
	return slice
}
