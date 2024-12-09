package rect

import "fmt"

func init() {
	fmt.Println("rect package is used ")
}

func New(l, b float32) *Rect { // using New is an idiomatic approach , to create a constructure
	return &Rect{L: l, B: b}
}
func NewDefault() *Rect {
	return &Rect{L: 0, B: 0}
}

type Rect struct { //exported
	L, B float32
}

func (r *Rect) Area() float32 {
	return r.L * r.B
}
func (r *Rect) Perimeter() float32 {
	return 2 * (r.L + r.B)
}

// func NewCuboid
