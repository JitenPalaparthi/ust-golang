package main

import "unsafe"

func main() {
	str1 := "Hello World!"
	println(str1)
	println("Size of str1:", unsafe.Sizeof(str1))
	println("Address of str1:", &[]byte(str1)[0])

	str1 = "Hello World"
	println(str1)
	println("Size of str1:", unsafe.Sizeof(str1))
	println("Address of str1:", &[]byte(str1)[0])

	str1 = "Hello UST Global minds!"
	println(str1)
	println("Size of str1:", unsafe.Sizeof(str1))
	println("Address of str1:", &[]byte(str1)[0])

	str1 = "Hello World, how are you doing, why the world is unrest. Lets call for the peace"
	println(str1)
	println("Size of str1:", unsafe.Sizeof(str1))
	println("Address of str1:", &[]byte(str1)[0])
}

// string is immutable?

/*string structure

pointer -> 8 bytes (on 64 bit machines)
length  -> 8 bytes   (on 64 bit machines)

*/
