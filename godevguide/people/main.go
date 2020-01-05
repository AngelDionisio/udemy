package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func (p *person) updateFirstName(firstName string) {
	p.firstName = firstName
}

func main() {
	bond := person{
		firstName: "James",
		lastName:  "Bond",
		contactInfo: contactInfo{
			email:   "bond@compass.com",
			zipCode: 10456,
		},
	}

	bond.updateFirstName("Jimmy")

	bond.print()
}
