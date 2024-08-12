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
	var agent person
	jim := person{
		firstName: "Jimbo",
		lastName:  "Clash",
		contact: contactInfo{
			email:   "juum@bo.com",
			zipCode: 12322,
		}}

	agent.firstName = "Jason"
	agent.lastName = "Baktus"
	agent.contact.email = "Jason.b@murder.com"

	fmt.Println(agent)
	fmt.Printf("%+v", agent)
	fmt.Printf("%+v", jim)
}
