package main

import (
	"errors"
	"fmt"
)

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
	fmt.Println("area of rect r1:", r1.A)
	fmt.Println("perimeter of rect r1:", r1.P)

	//(&r1).Area() no need to call like this .. just call r1.

	r2 := &Rectangle{L: 12.123, B: 123.123} // shorthand declararion
	a3 := r2.Area()
	p3 := r2.Perimeter()
	fmt.Println("Area of Rect r2:", a3)
	fmt.Println("Perimeter of Rect r2:", p3)
	fmt.Println("area of rect r2:", r2.A)
	fmt.Println("perimeter of rect r2:", r2.P)

	r3 := new(Rectangle)
	r3.L = 123.23
	r3.B = 12.23
	a4, p4 := r3.Area(), r3.Perimeter()
	fmt.Println("r3 area:", a4, "r3 perimeter:", p4)

	//r4 := &Rectangle{L: 12.123, B: 123.123}

	var r4 *Rectangle // r4 is nil, declared but not instantiated
	//(&r4).AreaAndPerimeter()
	a5, p5, err := r4.AreaAndPerimeter()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("r4 area:", a5, "r4 perimeter:", p5)
	}

	var r5 *Rectangle // r4 is nil, declared but not instantiated
	//r5 = new(Rectangle)
	r5 = &Rectangle{L: 12.23, B: 23}
	a6, p6, err := r5.AreaAndPerimeter()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("r5 area:", a6, "r5 perimeter:", p6)
	}

}

type Rectangle struct {
	L    float32 // 0
	B    float32 // 0
	A, P float32
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

func (r *Rectangle) Area() (area float32) {
	if r != nil {
		r.A = r.L * r.B
		return r.A
	}
	return 0.0
}

func (r *Rectangle) Perimeter() float32 {
	if r != nil {
		r.P = 2 * (r.L + r.B)
		return r.P
	}
	return 0
}

func (r *Rectangle) AreaAndPerimeter() (a float32, p float32, err error) {
	if r == nil {
		return 0, 0, errors.New("nil rectangle receiver")
	}
	return r.Area(), r.Perimeter(), nil
	//return r.L * r.B, 2 * (r.L + r.B)
}

// methods to be attached to the User defined types
// methods contain receivers
// receivers are for types(user defined)
