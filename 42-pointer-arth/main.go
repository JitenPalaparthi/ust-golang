package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	arr1 := []int8{10, 20, 30, 40, 50}
	var arrunitptr uintptr = uintptr(unsafe.Pointer(&arr1[0]))

	for i := 0; i < len(arr1); i++ {
		valuePtr := (*int8)(unsafe.Pointer(arrunitptr))
		println(*valuePtr)
		arrunitptr += unsafe.Sizeof(int8(0))
	}

	str1 := "Hello World" // pointer length

	struintptr := uintptr(unsafe.Pointer(&str1))
	struintptr += 8

	valuePtr := (*int)(unsafe.Pointer(struintptr))
	println("Length:", len(str1))
	println("length", *valuePtr)

	strheader := (*reflect.StringHeader)(unsafe.Pointer(&str1))

	println(strheader.Data)
	println(strheader.Len)

	strPtr := strheader.Data

	for i := 0; i < len(str1); i++ {
		btPtr := (*byte)(unsafe.Pointer(strPtr))
		println(string(*btPtr))
		strPtr += 1
	}

	slice1 := []int{10, 20, 30, 40, 50}
	/*
		ptr -> 8 bytes
		len -> 8 bytes
		cap -> 8 bytes
	*/

	sliceUintptr := uintptr(unsafe.Pointer(&slice1))

	structSlice1 := (*int)(unsafe.Pointer(sliceUintptr))
	println("Slice header pointer:", *structSlice1)

	fmt.Printf("address of slice:%d\n", &slice1[0])

	sliceUintptr += 8
	structSlice1 = (*int)(unsafe.Pointer(sliceUintptr))
	println("Slice Length:", *structSlice1)

	sliceUintptr += 8
	structSlice1 = (*int)(unsafe.Pointer(sliceUintptr))
	println("Slice Capasity:", *structSlice1)

	sliceUintptr = uintptr(unsafe.Pointer(&slice1))

	structSlice2 := (*[3]int)(unsafe.Pointer(sliceUintptr))
	println("Slice header pointer:", *structSlice1)

	for i := 0; i < len(structSlice2); i++ {
		println(structSlice2[i])
	}

}
