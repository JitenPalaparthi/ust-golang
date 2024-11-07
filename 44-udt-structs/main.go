package main

import "fmt"

func main() {
	r1 := Rectangle{L: 12.123, B: 123.123} // shorthand declararion

	a1 := Area(r1)
	fmt.Println("Area of Rect r1:", a1)
	p1 := Perimeter(r1)
	fmt.Println("Perimeter of Rect r1:", p1)

	a2 := r1.Area()
	p2 := r1.Perimeter()
	fmt.Println("Area of Rect r1:", a2)
	fmt.Println("Perimeter of Rect r1:", p2)

}

type Rectangle struct {
	L float32 // 0
	B float32 // 0
}

// Area and Perimeter are just functions and Rect variable is used as a param and called

func Area(r Rectangle) (area float32) {
	area = r.L * r.B
	return area
}

func Perimeter(r Rectangle) float32 {
	p := 2 * (r.L + r.B)
	return p
}

func (r Rectangle) Area() (area float32) {
	area = r.L * r.B
	return area
}

func (r Rectangle) Perimeter() float32 {
	p := 2 * (r.L + r.B)
	return p
}

// methods to be attached to the User defined types
// methods contain receivers
// receivers are for types(user defined)
