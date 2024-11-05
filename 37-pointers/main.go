package main

func main() {

	// var ptr1 *int // the zero value of a pointer is nil

	// // ptr1 = &100 // This is posible in rust but in go

	// var num1 int = 1231

	// ptr1 = &num1

	// there is a allocator using a built in function called new

	ptr2 := new(int) // allocate some memory which is of size int and give a zero value to the memory

	//1. ptr2 is a pointer
	//2. it is not a nil pointer
	//3. it has a zero value
	//4. it can be dereferenced since it is not nil

	println("address of ptr2:", ptr2)
	println("value of ptr2:", *ptr2)

	*ptr2 = (2 * 2) + (3 / 3) + 100 - 2

	// var ptr3 *int // nil
	var ptr3 *string = new(string)
	*ptr3 = "Hello world"
	println("address of ptr3:", ptr3)
	println("value of ptr3:", *ptr3)

	var ptr4 **string = &ptr3
	var ptr5 ***string = &ptr4

	println("deref of ptr5:", ***ptr5)

}

// if you create a pointer,
// the pointer need some address which is from another variable
//
