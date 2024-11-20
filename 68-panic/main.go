package main

import "fmt"

func main() {
	defer func() {
		println("End of Main")
	}()

	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println(r)
			}
		}()
		slice := []int{12, 14, 43, 34, 55, 5, 6, 3, 98, 65, 45, 78, 32}
		defer println()
		for i := 0; i <= len(slice); i++ {
			print(slice[i], " ")
		}
	}()

	println("can I be executed after panic?")

	var ptr *int
	*ptr = 100 // ? is it an error ? what is it?
	var n = 0
	println(100 / n)

}
