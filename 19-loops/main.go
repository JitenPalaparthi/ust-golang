package main

func main() {

	println("breaks only the inner loop")
	for i := 1; i <= 5; i++ {
		for j := 5; j >= 1; j-- {
			println("i:", i, "j:", j)
			if i == j {
				break
			}
		}
	}

	println("breaks outer loop as well")
outer:
	for i := 1; i <= 5; i++ {
		for j := 5; j >= 1; j-- {
			println("i:", i, "j:", j)
			if i == j {
				break outer
			}
		}
	}
}
