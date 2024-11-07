package main

import "fmt"

func main() {

	var m1 struct {
		id   int
		name string
	}

	m1.id = 123
	m1.name = "12312"

	m2 := struct {
		id   int
		name string
	}{
		100, "Jiten",
	}
	fmt.Println("one way m1:", m1)
	fmt.Println("one way m2:", m2)
}
