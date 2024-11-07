package main

import (
	"fmt"
)

func main() {

	p1 := &Person{id: 12, name: "Jiten", email: "jitenp@outlook.com", address: struct {
		doorno, city, pincode string
	}{doorno: "1-22", city: "bangalore", pincode: "560086"}}
	fmt.Println(p1)
	fmt.Println("Pincode:", p1.address.pincode)

	c1 := &Company{Id: 123, Name: "Abcd corp", Website: "www.abcd.io", Address: Address{DoorNo: "123/123", City: "hyd", Pincode: "542322"}}
	fmt.Println(c1)
	fmt.Println("City:", c1.Address.City)
	c1.Address.PrintAddress()

	s1 := &Student{RegNo: "123ab", Name: "Jiten", Email: "JitenP@Outlook.com", Contact: "9618558500", Address: &Address{DoorNo: "123/123", City: "hyd", Pincode: "542322", Status: "active"}}
	fmt.Println("City:", s1.City)
	fmt.Println("Status of student:", s1.Status)
	fmt.Println("Status of address:", s1.Address.Status)

	s1.PrintAddress()

	s := new(Some)
	s.Who("Jiten")

	(&Some{}).Who("Raj") // just calling without creating a func

	sq := (&Number{100}).Square()
	fmt.Println("square:", sq)
}

type Person struct {
	id          int
	name, email string
	address     struct { // embedded
		doorno, city, pincode string
	}
}

type Company struct {
	Id      int
	Name    string
	Website string
	//Addr    Address
	Address Address // okay to give same name
}

type Address struct {
	DoorNo, City, Pincode, Status string
}

func (a *Address) PrintAddress() {
	fmt.Println("DoorNo:", a.DoorNo, "City:", a.City, "Pincode:", a.Pincode, "Status:", a.Status)
}

// promoted

type Student struct {
	RegNo, Name, Email, Contact, Status string
	*Address                            //promoted fields
	// Address  // promoted field can be a pointer or a normal one
}

type Some struct{}

func (s *Some) Who(name string) {
	fmt.Println("Who:", name)
}

type Number struct {
	N int
}

func (n *Number) Square() int {
	return n.N * n.N
}
