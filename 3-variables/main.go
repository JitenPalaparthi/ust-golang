package main

import "fmt"

func main() {
	var num1 int = 12312
	println("num1:", num1)

	var num2 uint8 = 123

	var (
		num3, num4 uint64  = 1232321, 1232232
		float1     float32 = 1232.23
		float2     float64 = 12321.1231123
		ok1        bool    = true
		str1       string  = "UST Global"
	)

	var num5 = 123213 // automatically type is inferred

	var age1 = 39

	var (
		ok2 = true

		str2 = "Hello there"

		float3 = 13232.12321312
	)

	var ( // value is inferred based on type of the variable
		ok3 bool // false

		str3 string // empty string

		float4 float32 // zero -> 0.0
	)

	// shorthand declaration
	num6 := 12312
	float5 := 123.123
	ok4 := true
	str4 := "Hello World"

	println(num1, num2, num3, num4, num5, float1, float2, float3, float4, ok1, ok2, ok3, str1, str2, str3, age1)
	fmt.Printf("float1:%.2f float2:%v str1:%v ok1:%v\n", float1, float2, str1, ok1)
	println(num6, float5, ok4, str4)
}

/*
numbers
	int, uint, int8,int16,int32,int64,uint8,uint16,uint32,uint63,uintptr,float32,float64
	rune,byte, complex
boolean
	bool
strings
	string
interface
	interface{} or any
*/

// type and value inference
