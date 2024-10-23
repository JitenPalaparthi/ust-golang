package main

import (
	"fmt"
	"strconv"
)

func main() {

	var str1 string = "12321"

	num1, err := strconv.Atoi(str1)

	fmt.Println(num1, err)

	fmt.Printf("num1:%v type: %T\n", num1, num1)

	num2 := 3132131

	str2 := strconv.Itoa(num2)

	fmt.Printf("str2:%v type: %T\n", str2, str2)

	str3 := fmt.Sprint(num2)
	fmt.Printf("str3:%v type: %T\n", str3, str3)

	var (
		ok1    = true
		float1 = 12312.1231
		char1  = 'A'
	)

	str4 := fmt.Sprintf("%v %v %v", ok1, float1, char1) // formatter
	fmt.Printf("str4:%v type: %T\n", str4, str4)

}
