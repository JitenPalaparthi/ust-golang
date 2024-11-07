package main

import (
	"fmt"
	"unsafe"
)

func main() {

	var num1 int = 100
	var num2 int = 200

	//var uintptr1 uintptr =  &num1 // cant do that

	var uintptr1 uintptr = uintptr(unsafe.Pointer(&num1))
	fmt.Println("address of:", uintptr1)
	fmt.Printf("address %X\n", uintptr1)
	fmt.Printf("address %d\n", &num1)

	uintptr1 += 8
	fmt.Println("address of:", uintptr1)

	ptr1 := (*int)(unsafe.Pointer(uintptr1))

	println(num2, *ptr1)

	//var ptr1 *int = &num1

	//ptr1 += 8 // why this cant be done?

	arr1 := []int{10, 20, 30, 40}

	unitptr1 := uintptr(unsafe.Pointer(&arr1[0]))

	arrPtr1 := (*int)(unsafe.Pointer(unitptr1))
	println(*arrPtr1)

	unitptr1 += 8 // this is exactly pointer arthemetic here

	arrptr2 := (*int)(unsafe.Pointer(unitptr1))
	println(*arrptr2)

}
