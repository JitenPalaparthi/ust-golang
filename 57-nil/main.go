package main

import (
	"fmt"
)

func main() {

	var str1 string
	// if str1 == nil {
	// }
	fmt.Println(`str1 is not nil, but empty string ""`, str1)

	var ptr1 *int
	fmt.Println(ptr1)

	var myfunc func(string) = Greet
	if myfunc != nil {
		fmt.Println(myfunc)
		myfunc("How are you doing")
	}

	var s1 S1 = S1{num: 100}
	var s2 S1 // zero values are assigned to the fields in the struct
	// if s1 == nil {

	// }
	fmt.Println(s1, s2)

	var any1 any = nil
	fmt.Println(any1)

	var num1 int = 100
	var num2 int16 = 12312
	var ptr2 *int = &num1
	fmt.Println(ptr2, *ptr2)
	ptr2 = nil
	fmt.Println(ptr2)
	num3 := int(num2)
	ptr2 = &num3

	var slice1 []int
	fmt.Println(slice1)
	slice1 = append(slice1, 100)
	fmt.Println(slice1)

	var map1 map[int]int
	//map1[1] = 1
	fmt.Println(map1)

	var a1 any

	err := SetValue(a1, 100)

	fmt.Println(err.Error())
}

func Greet(msg string) {
	fmt.Println("Hello World", msg)
}

type S1 struct {
	num int
	ok  bool
}

func SetValue(a any, v int) error {
	if a == nil {
		m := &MyError{Msg: "any variable a is nil", Code: 1001}
		return m
	}

	a = v
	return nil
}

type MyError struct {
	Msg  string
	Code int
}

func (m *MyError) Error() string {
	return fmt.Sprintf("Message:%s Code:%d", m.Msg, m.Code)
}
