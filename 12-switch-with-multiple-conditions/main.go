package main

func main() {

	num := -25

	switch { // empty switch
	case (num >= 0 && num <= 50) && true && (true || false) && (true || !true):
		// true && ... are just to mimic the very big condition.
		// just to mimic , in case can write any kind of logic.
		ok := true
		switch ok {
		case true:
			println(num, " is between 0 and 50")
		case false:

		default:
		}

	case num > 50 && num <= 100:
		{
			println(num, "is above 50 and lessthan and equal to 100")
		}
	case num > 100:
		{
			println(num, "is above 100")
		}
	default:
		println(num, "is a negative number")
	}

	char := 'A'

	switch char {

	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		println(string(char), "is vovel")
	default:
		println(string(char), "is either consonent or a special character")
	}
}

// break is automatically enabled in go
