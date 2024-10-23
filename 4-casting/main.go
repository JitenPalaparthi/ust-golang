package main

import "fmt"

type (
	char    = int32
	integer = int
)

func main() {

	// short a = 10; // 2 bytes
	// long b = a; // 8 bytes

	var num1 uint8 = 230

	//var num2 uint64 = num1 // does not work

	var num2 uint64 = uint64(num1) // type casting

	//ok1 := true

	//var num3 uint8 = uint8(ok1) // cant do ..bcz ok1 is bool. bool cannot be casted to number

	var float1 float32 = 123.123

	var num3 int = int(float1)

	var char1 rune = 'A' // alias for int32

	var char2 int32 = 'B'

	var char3 char = '你'

	var char4 char = 65

	var charcode int = 20320

	println(num1, num2, num3)

	println(char1, char2, char3, char4)

	println(string(char1), string(char2), string(char3), string(char4))
	println(string(charcode))

	var num4 int = 123123123 //111010101101011010110110011

	fmt.Printf("binary of num4 %b\n", num4)

	var num5 uint8 = uint8(num4) // 10110011 ->179

	var num6 uint16 = uint16(num4) // 1011010110110011 --> 46515

	println(num1, num2, num3, num4, num5, num6)

}

// type casting
// go does not suppor implicit type casting
