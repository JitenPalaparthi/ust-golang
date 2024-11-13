package circle

import "fmt"

const pi = 3.14 // P should be in upper case inorder to be accessed from main.go
// here pi is accessed inside the same pacakge so okay to give pi in lowercase
// but as a coding standards  constants are ALL UPPERCASES
var R float32

func Perimeter() float32 {
	return 2 * pi * R
}

func Area() float32 {
	return pi * R * R
}

func init() {
	fmt.Println("circle package is used ")
}

func init() {
	fmt.Println("circle package is used, again calling init")
}

func init() {
	R = 1.0
}
