package main

func main() {
	defer greet()
	defer println("ending main-1") // f1 deferes the exeution to the end of the call stack
	defer println("end of main-2") // f2
	println("start of main")
}

/*
f2
f1
*/

func greet() {
	defer println("end of greet function")
	println("Hello gret function")
}
