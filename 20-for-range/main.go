package main

func main() {
	str1 := "Hello, World!" // what is the size of each character?
	str2 := "Hello, 世界😀!"   // what is the size of each character?

	char := 64000
	println(string(char))

	char1 := "😀"
	char2 := "世"
	char3 := "A"
	println(len(char1), len(char2), len(char3))

	// unicode chars are not eactly 4 bytes but they are from 1-4 bytes
	// using normal for loop

	for i := 0; i < len(str1); i++ {
		//byte of a string
		//print(str1[i], " ")
		print(string(str1[i]), " ")
	}
	println()
	for i := 0; i < len(str2); i++ {
		//byte of a string
		print(string(str2[i]), " ")
	}

	println()
	// range loop how characters are taken
	for i, v := range str2 {
		println("i:", i, "char-> ", string(v), "-->", v)
	}
}
