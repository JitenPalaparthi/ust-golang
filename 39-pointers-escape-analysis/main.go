package main

// var (
// 	ptr = new(int)
// )

var num2 int = 100

const (
	PI = 3.14
)

func main() {

	var num = 100

	// num = (num+100)*(num/2) + (num * 2) - num%10

	n1 := getSquare(num)
	println("n1:", n1)

	getSquareP(&num)
	println("num:", num)

	// num = 100

	p1 := getSquarePRS(num)
	println("p1:", *p1)
	*p1 = 100 * 200
	println("p1:", *p1)

	// p2 := getSquarePR(num)
	// println("p2:", *p2)

	// p3 := getSquarePRG(num)
	// println("p3:", *p3)
}

func getSquare(num int) int { // call by value
	r := num * num
	return r
}

func getSquareP(num *int) { // call by ref
	if num != nil {
		*num = *num * *num
	}
}

func getSquarePRS(num int) *int {
	num = num * num
	return &num // outliving the variable out of the scope
}

// func getSquarePR(num int) *int {
// 	ptr1 := new(int)
// 	*ptr1 = num * num
// 	return ptr1
// }

// func getSquarePRG(num int) *int {
// 	*ptr = num * num
// 	return ptr
// }

// // dangling pointers -> use a pointer which is already freed
// // nil pointer dereference --> a pointer is already freed and nil , you try to dereference a pointer
// // double free a pointer-> a pointer is already freed, again try to free a freed pointer
// // data races -> pointers work with addresses and you work with those pointers in concurrent code
// // memory leak -> create a pointer in side a scope and try to use it outside of the scope
