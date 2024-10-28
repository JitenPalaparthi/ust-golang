package main

func main() {

	var a, b = 10, 20

	var exp1 = (a * a) + (b * b) + 2*a*b

	var exp2 = a > b && true || false
	// false && true || false
	// false || false
	// false
	println(exp1, exp2)
}

// arthimetic
// logical
// comparision

// operator precedence

// 1. ()
// 2. *,/,%, &, +, -, !, ^, &^, <- (unary operators)
// 3. *, /, %, <<, >>, &, &^
// 4. +,-
// 5. ==, !=, <, <=, >, >=
// 6. <- (send operations on channel)
// 7. &&,||
// 8. =, +=, -=, *=, /=, %=, <<=, >>=, &=, &^=, `
