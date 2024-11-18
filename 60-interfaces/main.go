package main

import (
	"fmt"
	"shapes/shape"
	"shapes/shape/cuboid"
	"shapes/shape/rect"
	"shapes/shape/square"

	"math/rand"
)

func main() {

	slice1 := make([]shape.Shaper, 5)
	slice1[0] = &rect.Rect{L: 10.20, B: 12.23}
	slice1[1] = &cuboid.Cuboid{L: 10.20, B: 12.23, H: 13.45}
	slice1[2] = &rect.Rect{L: 16.20, B: 18.23}
	slice1[3] = &cuboid.Cuboid{L: 10.20, B: 12.23, H: 12.23}
	s1 := new(square.Square)
	*s1 = 123.123
	slice1[4] = s1

	// for _, v := range slice1 {
	// 	Shaper(v)
	// }

	// for i := 1; i <= 5; i++ {
	// 	index := rand.Intn(5)
	// 	Shaper(slice1[index])
	// }
	c := 0
	executed := false
out:
	if !executed {
		println("Jumping to in lable")
		goto in
	} else {
		goto finish
	}

in:
	{
		index := rand.Intn(5)
		Shaper(slice1[index])
		c++
		if c < 5 {
			goto in
		} else {
			executed = true
			goto out
		}
	}
finish:
	println("Done finished")

	var num1 = 100
	var ok = true
	var r1 = &rect.Rect{L: 123, B: 23}
	Print(num1)
	Print(ok)
	Print(r1)

}

func Shaper(shaper shape.Shaper) {
	fmt.Println("Area of", shaper.What(), ":", shaper.Area())
	fmt.Println("Perimeter of", shaper.What(), ":", shaper.Perimeter())
}

func Print(e shape.IEmpty) {
	fmt.Println(e)
}
