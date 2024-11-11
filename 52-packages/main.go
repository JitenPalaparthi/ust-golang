package main

import (
	"fmt"
	"shapes/shape"
	"shapes/shape/circle"
	"shapes/shape/rect"

	"shapes/shape/square"

	cuboid "shapes/shape/cubid" // cubid is purposely given to demonstrate the name of
	// the directory is not 100% necessarily the name of the package
)

func init() {
	fmt.Println("init in main")
}

func main() {

	shape.What()

	r1 := &rect.Rect{L: 123.123, B: 123.43}
	a1, p1 := r1.Area(), r1.Perimeter()
	fmt.Println("Area of r1:", a1, "Perimeter of r1:", p1)

	s1 := square.Square(25.25)
	a2, p2 := s1.Area(), s1.Perimeter()
	fmt.Println("Area of s1:", a2, "Perimeter of s1:", p2)

	c1 := cuboid.Cuboid{L: 123.2, B: 43.34, H: 35.76} // name of the directory and the name of the package not
	// necessarily same name but as a convenction same name is given
	fmt.Println("Area of c1:", c1.Area(), "Perimeter of c1:", c1.Perimeter())
	//c1 := &rect.Cuboid{l: 12.23, B: 123.23, H: 13.23} // cannot use cuboid

	circle.R = 12.123
	fmt.Println("Area of circle:", circle.Area(), "Perimeter of circle:", circle.Perimeter())

	//r2 := &rect.Rect{L: 34.54, B: 78.45}
	r2 := rect.New(12.34, 54.34)

	a4, p4 := r2.Area(), r2.Perimeter()
	fmt.Println("Area of r2:", a4, "Perimeter of r2:", p4)

	circle.R = 25.34
	fmt.Println("Area of circle:", circle.Area(), "Perimeter of circle:", circle.Perimeter())

}

// every package should have a directory
// idiomatically the name of the package is the name of the directory
// idiomatically package names are in lower case and does not contain _ etc..
// each go file must have a package
// there are no access modifiers or specifiers in go (go does not have public private etc)
// there are only exported or unexported
