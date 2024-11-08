package main

import (
	"fmt"
	"strconv"
)

type myint int // new type

func (m myint) ToString() string {
	return strconv.Itoa(int(m))
}

func (m myint) Square() int {
	return int(m * m)
}

// another type
type anotherMyint myint

func (m anotherMyint) Cube() int {
	return int(m * m * m)
}

func main() {
	var num1 int = 12312
	var mynum1 myint = 123123
	var anothermynum1 anotherMyint = 123321

	s1 := myint(num1).ToString()
	sq1 := myint(num1).Square()
	cb1 := anotherMyint(num1).Cube()

	fmt.Println("num1 ToString:", s1, "Square:", sq1, "Cube:", cb1)

	s2 := mynum1.ToString()
	sq2 := mynum1.Square()
	cb2 := anotherMyint(mynum1).Cube()

	fmt.Println("mynum1 ToString:", s2, "Square:", sq2, "Cube:", cb2)

	s3 := myint(anothermynum1).ToString()
	sq3 := myint(anothermynum1).Square()
	cb3 := anothermynum1.Cube()

	fmt.Println("anothermynum1 ToString:", s3, "Square:", sq3, "Cube:", cb3)

	var float1 float64 = float64(anothermynum1)
	fmt.Println("float can be type casted to anothermyint,myint and vise versa:", float1)
}
