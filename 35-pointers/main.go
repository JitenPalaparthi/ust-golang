package main

//var G int

func main() {
	var num1 int = 100
	var ptr1 *int = &num1
	println("direct: address of num1:", &num1, "value in num1:", num1) // dereference
	println("pointer:address of num1:", ptr1, "value in num1:", *ptr1) // dereference

	if ptr1 != nil {
		*ptr1 = *ptr1 + 100 // dereference a pointer with a new value by adding 100
	}
	println("num1:", num1)

	var ptr2 *int // nil
	if ptr2 == nil {
		println("nil pointer", ptr2)
	}

}

// null pointer (in go it is called nil pointer)
// dereferenceing a nil/null pointer
//
