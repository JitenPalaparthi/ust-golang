package main

import "fmt"

func main() {

	c1 := complex(12.12, 123.23) // complex128
	var r1, i1 float32 = 12.12, 123.23
	c2 := complex(r1, i1) // complex64 bcz r1 and i1 are float32
	c3 := 12.12 + 123.23i //(r+ni) this is also a complex number which is complex128
	var c4 complex128 = c1 + c3
	c5 := c1 + complex128(c2) // can do type casting
	c6 := c1 - c3
	c7 := c1 * complex128(c2)
	c8 := complex64(c1) / c2

	r2 := real(c8) // fetch real part from a complex number. since c8 is complex64 both real and imagary are float64
	i2 := imag(c8) // fetch imaginary part from a complex number

	r3 := real(c1) // fetch real part from a complex number. since c8 is complex64 both real and imagary are float64
	i3 := imag(c1) // fetch imaginary part from a complex number

	fmt.Println(c1, c2, c3, c4, c5, c6, c7, c8)
	fmt.Println(r2, i2, r3, i3)

}

// two types of complex numbers. complex64, complex128
