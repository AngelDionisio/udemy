// -----------------------------------------
// Unexported fields from an exported struct
// -----------------------------------------
package main

import (
	"fmt"

	// relative path to a physical location on disk - relative to GOPATH.
	"github.com/angeldionisio/udemy/golang/random/concepts/exporting/users"
)

func main() {
	// Creating a value of type User from the users package
	// since password is unexported, it cannot be compiled
	// yields the following error if uncommented:
	// unknown field 'password' in struct literal of type users.User
	user := users.User{
		Name: "Angel",
		ID:   975,

		// password: "xxxxxx",
	}

	fmt.Printf("%#v\n", user)

	// using a setter to modify an unimported property
	user.SetPassword("xxxxxxx")

	fmt.Printf("%#v\n", user)

}
