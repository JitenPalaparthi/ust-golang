package square

import "fmt"

func init() {
	fmt.Println("square package is used ")
}

type Square float32

func (s Square) Area() float32 {
	return float32(s * s)
}

func (s Square) Perimeter() float32 {
	return float32(4 * s)
}
