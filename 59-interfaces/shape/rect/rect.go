package rect

import "fmt"

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
