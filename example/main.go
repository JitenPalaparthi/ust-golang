package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	var i any = 42 // Assigning an integer to an `any` type

	// Inspect the internal structure of `any`
	iface := (*[2]uintptr)(unsafe.Pointer(&i))
	typeInfo := iface[0] // Type information
	dataPtr := iface[1]  // Pointer to actual data

	fmt.Printf("Type Info Address: %x\n", typeInfo)
	fmt.Printf("Data Pointer Address: %x\n", dataPtr)

	// Using reflection to get more details
	value := reflect.ValueOf(i)
	fmt.Printf("Concrete Type: %s\n", value.Type())
	fmt.Printf("Concrete Value: %v\n", value.Interface())

	var str string

	ptr2 := (*[2]int)(unsafe.Pointer(&str))

	fmt.Println(*ptr2, "type of", reflect.TypeOf(ptr2))

	var fn1 func() = func() {
		println("Hello World")
	}

	ptr3 := (*[10]int)(unsafe.Pointer(&fn1))

	fmt.Println(*ptr3, "type of", reflect.TypeOf(ptr3))

}
