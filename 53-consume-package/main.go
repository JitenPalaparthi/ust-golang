package main

import (
	"fmt"

	"github.com/JitenPalaparthi/ust-shapes/shape"
	"github.com/JitenPalaparthi/ust-shapes/shape/circle"
	"github.com/JitenPalaparthi/ust-shapes/shape/cuboid"
	"github.com/JitenPalaparthi/ust-shapes/shape/rect"
	"github.com/JitenPalaparthi/ust-shapes/shape/square"
)

func main() {

	shape.What()
	shape.Why()

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
