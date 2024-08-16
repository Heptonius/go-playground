package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	jim := person{
		firstName: "Jimbo",
		lastName:  "Clash",
		contact: contactInfo{
			email:   "juum@bo.com",
			zipCode: 12322,
		}}

	jimPointer := &jim
	jimPointer.updateFirstName("Jimmy")
	jim.print()
}

func (pPointer *person) updateFirstName(newFirstName string) {
	(*pPointer).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
