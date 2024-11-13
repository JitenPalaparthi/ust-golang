package main

import "fmt"

func main() {
	Greet()
	(Greeting{}).Greet("Hello UST Global minds")

	st1 := struct {
		F1 int
		F2 float32
	}{F1: 100, F2: 340.23}
	fmt.Println(st1)

	var st2 struct {
		F1 int
		F2 float32
	}
	st2.F1 = 100
	st2.F2 = 123.123

}

func Greet() {
	fmt.Println("Hello World")
}

type Greeting struct{} // empty no fields

func (g Greeting) Greet(msg string) {
	fmt.Println(msg)
}
