package main

import (
	"fmt"
	"os"
)

func main() {

	file, err := os.Open("go.mod")
	if err != nil {
		fatal(err.Error())
	}
	defer file.Close()

	println("start of main")

	//read := true
	for {
		buf := make([]byte, 10)
		_, err := file.Read(buf)
		fmt.Print(string(buf))
		if err != nil {
			fmt.Println("Error", err.Error())
			break
		}
	}

}

func fatal(msg ...any) {
	fmt.Println(msg...)
	os.Exit(1)
}
