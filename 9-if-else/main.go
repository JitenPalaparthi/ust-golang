package main

func main() {

	// if even or not

	num := 102

	if num%2 == 0 {
		println(num, "is even number")
	} else {
		println(num, "is odd number")
	}

	// same thing in a different syntax , along with creation of variable

	//num := 102

	if num2 := 101; num2%2 == 0 {
		println(num2, "is even number")
	} else {
		println(num2, "is odd number")
	}

	//println(num2) , cant use it out of if else block

}
