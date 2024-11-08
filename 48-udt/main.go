package main

import (
	"fmt"
	"strconv"
)

type integer = int // this is not creating a new type. It is just aliasing

type myint int // new type

type anotherMyint myint

type NumStruct struct {
	Num int
}

func (ns *NumStruct) ToInt() int {
	return ns.Num
}

func (m myint) ToString() string {
	return strconv.Itoa(int(m))
}

func (m myint) Square() int {
	return int(m * m)
}

func (m anotherMyint) Cube() int {
	return int(m * m * m)
}

func main() {
	var myint1 myint = 100
	//myint2 := myint(100)
	s1 := myint1.ToString()
	s2 := myint1.Square()
	fmt.Println("To String of myint1:", s1)
	fmt.Println("Square of myint1:", s2)

	var num1 int = 25

	var num2 uint8 = 25

	var float1 float32 = 123.34

	fmt.Println(myint(num1).ToString())
	fmt.Println(myint(num2).Square())
	fmt.Println(myint(float1).ToString())

	// var ok1 bool

	//myint(ok1).ToString() // cant do it as with int to bool cant cast

	var ns1 NumStruct = NumStruct{Num: 100}

	//var num3 int = int(ns1) // cant to this
	var num3 anotherMyint = anotherMyint(myint(ns1.ToInt()))

	//strconv.Itoa(100)
	println(num3)

}
