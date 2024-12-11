package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	password := "jiten"
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(bytes))
	//$2a$10$wLtslY5qIDfOUprXAxluxuh0m13tTiG9M2fVeXTZbK0ZTyv59UDue

	// to check password

	err = bcrypt.CompareHashAndPassword(bytes, []byte("jiten"))
	if err == nil {
		fmt.Println("valid password")
	} else {
		fmt.Println("invalid password")
	}
}
