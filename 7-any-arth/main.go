package main

import (
	"fmt"
	"strconv"
)

func main() {
	var (
		num1   uint8       = 180
		float1 float32     = 123
		any2   any         = 213.123
		any3   interface{} = 1233
		float2 float64     = 123.1231
		str1   string      = "123"
		str2   string      = "12323.123"
		ok1    bool        = true // if true add 1 else add 0
		sum    float64
	)

	f1, _ := strconv.ParseFloat(str1, 64) // strconv

	f2, err := strconv.ParseFloat(str2, 64) // strconv
	if err != nil {
		fmt.Println(err)
		return
	}

	b1 := 0.0
	if ok1 {
		b1 = 1.0
	} else {
		b1 = 0.0
	}

	sum = float64(num1) + float64(float1) + any2.(float64) + float64(any3.(int)) + float2 + f1 + f2 + b1

	fmt.Printf("sum:%.2f\n", sum)
}
