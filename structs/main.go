package main

import "fmt"

type person struct {
	FirstName   string
	LastName    string
	ContactInfo contactInfo
}

type contactInfo struct {
	Email   string
	ZipCode string
}

func main() {
	alex := person{
		FirstName: "Alex", LastName: "Anderson",
		ContactInfo: contactInfo{
			Email:   "my.sisler.email@gmail.com",
			ZipCode: "12345",
		},
	}

	alex.print()
	alex.updateName("Jim")
	alex.print()
}

func (pointerToPerson *person) updateName(updatedFirstname string) {
	(*pointerToPerson).FirstName = updatedFirstname
}

func (p person) print() {
	fmt.Println("")
	fmt.Printf("%+v", p)
}
