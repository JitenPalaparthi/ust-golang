package main

func main() {
	func() {
		println("Hello Welcome to Main function")
	}()
	defer func() {
		println("End of main func")
	}()
	fullName("Jiten", "P")
	fullName("Rajesh", "")
	func() {
		for i := 1; i <= 10; i++ {
			if i%2 == 0 {
				print(i, " ")
			}
		}
		println()
	}()
}

func fullName(firstName, lastName string) {
	func() {
		println("Welcome to fullName function")
	}()
	defer func() {
		println("End of fullname function")
	}()

	if firstName == "" {
		panic("First Name cannot be empty")
	}
	if lastName == "" {
		panic("Last Name cannot be empty")
	}
	println("Full Name:", firstName+" "+lastName)
}

// by default, panic panics the whole call stack.
// it unwinds everything after the panic
// but if there are any defers in the scope of panic , defers are executed before the panic
