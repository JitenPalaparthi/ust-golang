package main

import (
	"fmt"

	"math/rand"
)

func main() {
	var slice1 []int // a slice is defined but not instantiated, so it is a nil slice

	if slice1 == nil {
		slice1 = make([]int, 5) // while instantiating, if you dont give the capasity, the length become the capasity
	}

	fmt.Println("zero values:", slice1) // developer friendly print

	// short hand declaration of slice

	slice2 := make([]bool, 3)           // define and instantiate
	fmt.Println("zero values:", slice2) // developer friendly print

	// another way to declare a slice , directly assigning values
	slice3 := []string{"Jiten", "Rahim", "Rajesh", "John", "Virat", "Bindu", "Priya"}

	fmt.Println("slice3:", slice3)

	for index := range slice1 {
		slice1[index] = rand.Intn(100)
	}
	slice4 := make([]int, 5) // length and capacity are 5
	for i := 0; i < len(slice1); i++ {
		slice4[i] = rand.Intn(100)
	}

	slice5 := []int{40, 64, 19, 37, 99} // no length to be give []int

	fmt.Println("slice1:", slice1)
	fmt.Println("slice4:", slice4)
	fmt.Println("slice5:", slice5)

}

// slice is a reference type
// slice can grow or shirnk
// does slice store in heap? it depends , compiler decides, but we can anticipate it
// an array can be converted to slice using a notation
// slice can be nil
// to create a slice , generally make built in function is used
// a slice has a header(internally),similar to string
// there is also zero value(s) for the slice.Value inferred based on the type
