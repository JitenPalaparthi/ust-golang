package main

import (
	"fmt"
	"shapes/shape"
	"shapes/shape/rect"
	"shapes/shape/square"
)

func main() {

	// static dispatch
	// compiler is aware of on what object Area and Perimeter are callecd
	// during compiler time
	r1 := &rect.Rect{L: 123.23, B: 23.34}

	var s1 square.Square = 123.123

	a1 := (*r1).Area() // (*r1) not required , call directly
	p1 := r1.Perimeter()

	fmt.Println("Area of rect r1 :", a1)
	fmt.Println("Perimeter of rect r1 :", p1)

	// end of static dispatch

	// dynamic dispatch

	//var shaper shape.Shaper = r1
	Area(r1)
	//shaper = s1
	Area(s1)

}

func Area(shaper shape.Shaper) {
	fmt.Println("Area:", shaper.Area())
}

/* Intarface Method Table (v-table in other programming)
r1 --> 0x1231223
s1 -> 0x12312af
*/
