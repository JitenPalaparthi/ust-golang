package main

import "fmt"

func main() {

	r1 := &Rect{L: 10.23, B: 12.23}
	var s1 Square = 123.123
	s2 := Square(123.23)

	var shaper Shaper

	if shaper == nil { // can check interface is nil or not
		fmt.Println("Yes shapre is nil")
		shaper = r1 // extracted interface from a concrete object rect r1
	}

	a1 := shaper.Area()
	p1 := shaper.Perimeter()

	fmt.Println("Area a1:", a1)
	fmt.Println("Perimeter p1:", p1)

	shaper = s1

	a2 := shaper.Area()
	p2 := shaper.Perimeter()

	fmt.Println("Area a2:", a2)
	fmt.Println("Perimeter p2:", p2)

	shaper = s2

	a3 := shaper.Area()
	p3 := shaper.Perimeter()

	fmt.Println("Area a3:", a3)
	fmt.Println("Perimeter p3:", p3)

}

type Shaper interface {
	// Area() float32
	// Perimeter() float32
	IArea
	IPerimeter
}

type IArea interface {
	Area() float32
}
type IPerimeter interface {
	Perimeter() float32
}

type Rect struct {
	L, B float32
}

func (r *Rect) What() {
	fmt.Println("It is rect")
}

func (r *Rect) Area() float32 {
	return r.L * r.B
}
func (r *Rect) Perimeter() float32 {
	return 2 * (r.L + r.B)
}

type Square float32

func (s Square) Area() float32 {
	return float32(s * s)
}

func (s Square) Perimeter() float32 {
	return float32(4 * s)
}

// var any1 any // header -> ptr, type
// var str1 string // header -> ptr, len
