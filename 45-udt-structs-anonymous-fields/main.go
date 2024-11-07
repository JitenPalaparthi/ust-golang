package main

import "fmt"

func main() {

	c1 := ColorCode{int: 123, string: "red"}
	c1.Print()
	fmt.Println(c1.Get())
}

// anonymous fields in a struct
// code of the color
// name of the color
// type ColorCode struct {
// 	Code int
// 	Name string
// }

// type mystring = string // this is an alias for existing string and same as string
type ColorCode struct {
	int
	string
	//mystring // since the name here is different than string , it can be used as an another anonymous field
}

func (c *ColorCode) Print() {
	fmt.Println("Colour code:", c.int)
	fmt.Println("Colour Name:", c.string)
}

func (c *ColorCode) Get() (code int, name string) {
	return c.int, c.string
}
