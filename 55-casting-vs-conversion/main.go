package main

import (
	"fmt"
	"strconv"
)

type cint int

func (ci cint) ToString() string {
	return fmt.Sprint(ci)
}

func (ci cint) Square() cint {
	return ci * ci
}

func main() {

	var num1 int = 25

	var cnum1 cint = cint(num1) // type casting int to cint

	str1 := cnum1.ToString()
	sq1 := cnum1.Square()
	fmt.Println(str1, sq1)

	// --- conversion

	var str2 = "12.34a"

	float1, err := strconv.ParseFloat(str2, 64)

	fmt.Println(float1, err)

	float2, err := cstring(str2).SureParseFloat()
	fmt.Println(float2, err)

	var str3 cstring = "1xy12.34343abcdef1"

	float3, err := str3.SureParseFloat()
	fmt.Println(float3, err)

}

type cstring string

func (cs cstring) SureParseFloat() (float64, error) {
	str1 := ""
	var count uint8
	for _, char := range cs {
		switch char {
		case 48, 49, 50, 51, 52, 53, 54, 55, 56, 57:
			str1 += string(char)
		case 46:
			if count == 0 {
				str1 += string(char)
			}
			count++
		}
	}
	if str1 == "" {
		return 0, nil
	}
	return strconv.ParseFloat(str1, 64)

}
