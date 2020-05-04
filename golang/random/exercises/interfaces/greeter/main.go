package main

import (
	"fmt"
)

// User base person struct
type User struct {
	FirstName string
	LastName  string
}

// Greet prints User info, implements the  Greeter interface
func (u *User) Greet() string {
	return fmt.Sprintf("%s, %s", u.LastName, u.FirstName)
}

// Customer is a type of User
type Customer struct {
	CustomerID int
	FullName   string
}

// Greet prints Customer details, implements the Greeter interface
func (c *Customer) Greet() string {
	return fmt.Sprintf("customerID: %d, %s", c.CustomerID, c.FullName)
}

// Greeter interface definition
type Greeter interface {
	Greet() string
}

// GreetPerson executes the Greet function of Greeter types
func GreetPerson(g Greeter) string {
	return fmt.Sprintf("Dear Greeter: %s", g.Greet())
}

func main() {
	u := &User{
		FirstName: "James",
		LastName:  "Bond",
	}
	fmt.Println(u.Greet())
	fmt.Println(GreetPerson(u))

	c := &Customer{
		CustomerID: 7401,
		FullName:   "Miss MoneyPenny",
	}
	fmt.Println(GreetPerson(c))
}
