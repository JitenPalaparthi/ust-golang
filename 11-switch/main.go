package main

import "math/rand"

func main() {

	switch day := rand.Intn(10); day {
	case 1:
		println("sunday")
	case 2:
		println("monday")
	case 3:
		println("tuesday")
	case 4:
		{
			println("wednesday")
		}
	case 5:
		println("thursday")
	case 6:
		println("friday")
	case 7:
		println("saturday")
	default:
		println("noday->", day)
	}

}

// break is automatically enabled in go
