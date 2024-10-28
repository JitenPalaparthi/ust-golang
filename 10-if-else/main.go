package main

func main() {

	var (
		age    uint8 = 38
		gender rune  = 'M'
	)

	// true && (flase || true)
	// true && true
	// true

	ok1 := age >= 21 && (gender == 'm' || gender == 'M')
	if age >= 18 && (gender == 'f' || gender == 'F') {
		println("she is eligible for marraige")
	} else if ok1 {
		println("he is eligible for marriage")
	} else {
		println("not eligible for marriage")
	}
}
