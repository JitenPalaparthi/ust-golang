package main

import "math/rand"

func main() {

	println("print 1- 10 numbers")

	for i := 1; i <= 10; i++ {
		print(i, " ")
	}
	println()

	println("for loop like a while loop")

	i := 10 // init

	for i <= 20 { // like a while loop only with a condition
		print(i, " ")
		i++ // post condition
	}
	println()

	j, k := 1, rand.Intn(20)
	println("unconditional loop.Loop until it reaches the random number:", k)
	for { // a loop without any condition
		if j > k {
			break // this breaks the loop
		}
		print(j, " ")
		j++
	}
	println()

	println("even numbers")
	j = 2
	for ; j <= 10; j += 2 {
		print(j, " ")
	}

	println()

	println("odd numbers")

	for j := 1; j <= 10; {
		print(j, " ")
		j += 2
	}

	println()
}

// go has only for loop and for - range loop
