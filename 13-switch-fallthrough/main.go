package main

func main() {

	num := 24

	switch {
	case num%8 == 0:
		println(num, "is divisible by 8")
		fallthrough
	case num%4 == 0:
		println(num, "is divisible by 4")
		fallthrough
	case num%2 == 0:
		println(num, "is divisible by 2")
	default:
		println("not divisible by 8 or 4 or 2")
	}
	println("another fallthru with false negative")

	switch num := 26; { // variables can be declared in switch
	case num%2 == 0:
		println(num, "is divisible by 2")
		fallthrough
	case num%4 == 0:
		println(num, "is divisible by 4")
		fallthrough
	case num%8 == 0:
		println(num, "is divisible by 8")
	default:
		println("not divisible by 8 or 4 or 2")
	}
}
