package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	_ "time"
)

func main() {
	greet1()
	greet2("Hello UST Folks!")
	str3 := greet3("Hello UST Folks!")
	fmt.Println(str3)
	a1 := areaOfRect(12.23, 23.23)
	p1 := perimeterOfRect(12.13, 23.23)
	fmt.Println("Area of Rect:", a1, "Perimeter of Rect:", p1)

	a2, p2 := areaAndPerimeterOfRect1(12.23, 23.23)
	fmt.Println("Area of Rect:", a2, "Perimeter of Rect:", p2)

	a3, _ := areaAndPerimeterOfRect1(12.23, 23.23) // dont want perimeter from this function so give blank identifier
	fmt.Println("Area of Rect:", a3)

	_, p3 := areaAndPerimeterOfRect1(12.23, 23.23)
	fmt.Println("Perimeter of Rect:", p3)

	a, _, c := 10, 20, 30
	println(a, c)

	var str1 = "123"

	var num int

	num, _ = strconv.Atoi(str1) // you can omit error but not a correct approach
	// if err != nil { // is there is an error
	// 	println(err.Error())
	// 	return
	// }

	println(num)

	a4, p4, err4 := areaAndPerimeterOfRect2(12.23, 23.23)
	if err4 != nil { // is there is an error
		println(err4.Error())
		return
	}
	fmt.Println("Area of Rect:", a4, "Perimeter of Rect:", p4)

	// not a good approach to return err as a first return type.
	// always return err as the last return value
	err5, a5, p5 := areaAndPerimeterOfRect3(0, 23.23)
	if err5 != nil { // is there is an error
		println(err5.Error())
		return
	}
	fmt.Println("Area of Rect:", a5, "Perimeter of Rect:", p5)
	var str2 = "12313123s"

	var num2 int

	num2, err := strconv.Atoi(str2)
	if err != nil { // is there is an error
		println(err.Error())
		return
	}

	println(num2)
}

func greet1() {
	println("Hello UST Global")
}

func greet2(msg string) { // what is the size of the stack frame?
	println(msg)
}

func greet3(msg string) string {
	return strings.ToUpper(msg)
}

func areaOfRect(l float32, b float32) float32 { // unnamed return type
	a := l * b
	return a
}

func perimeterOfRect(l, b float32) (p float32) {
	p = 2 * (l + b)
	return p
}

func areaAndPerimeterOfRect1(l, b float32) (area float32, perimeters float32) {
	return areaOfRect(l, b), perimeterOfRect(l, b)
}

func areaAndPerimeterOfRect2(l, b float32) (area float32, perimeters float32, err error) {
	if l == 0 || b == 0 {
		return area, perimeters, errors.New("length(l) or bredth(b) is zero")
	}
	return areaOfRect(l, b), perimeterOfRect(l, b), nil
}

// technically you can return err type in any order
// but it is not idiomatic go approach
// err is a last return type everywhere in go
func areaAndPerimeterOfRect3(l, b float32) (err error, area float32, perimeters float32) {
	if l == 0 || b == 0 {
		return errors.New("length(l) or bredth(b) is zero"), area, perimeters
	}
	return nil, areaOfRect(l, b), perimeterOfRect(l, b)
}
