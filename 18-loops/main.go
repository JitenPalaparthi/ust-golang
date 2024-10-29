package main

import "math/rand"

func main() {

	for i, j := 1, 10; i <= j && true && (true || false); i, j = i+1, j-1 { // just simulating a big condition
		println("i:", i, "j", j)
	}

	println("continue")

	for i := 1; i <= 10; i++ {
		if i%2 == 1 {
			continue // continues the loop with next iteration by invoking the post condition
		}
		print(i, " ")
	}

	println("1-rand prime numbers")

	j := rand.Intn(100)
	if j >= 3 {
		print("1 2 3 ")
	}
	for i := 4; i <= j; i++ {
		c := 0
		for k := 2; k < i; k++ {
			if i%k == 0 {
				c++
				break
			}
		}
		if c < 1 {
			print(i, " ")
		}
	}

}
