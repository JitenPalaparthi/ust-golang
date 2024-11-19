package main

import "fmt"

func main() {

	str1 := new(string)
	*str1 = "Hello World!"

	// rstr := ""
	// for _, v := range str1 {
	// 	rstr = string(v) + rstr
	// }

	// fmt.Println(rstr)

	//c := 100 + 200/300%2

	for _, v := range *str1 {
		defer print(string(v))
	}
	println()

	num := 10
	ptr := &num

	defer fmt.Println("num with defer:", num, *ptr)
	num = 20
	fmt.Println("num normally:", num, *ptr)

}

// 5 mins of time
// take an array or slice
// print the elements of the array of slice in reverse using defer
