// Package users provides support for user management
package users

import (
	"log"
)

// User exported type represents information about an user
// It has t wo exported fields, Name, Email, and 1 unexported, password
type User struct {
	Name string
	ID   int

	password string
}

// SetPassword manages mutation of password field
func (u *User) SetPassword(pass string) {
	u.password = pass

	log.Printf("successfully updated password for userID: %v\n", u.ID)
}

// Notes:
/*
	if we have a type that is unimported, but has exported fields. If this struct is
	embedded into another struct, even though the struct itself is unimported, the exported
	fields can still be set. E.G:

	// user represents information about a user.
	// Unexported type with 2 exported fields.
	type user struct {
		Name string
		ID   int
	}

	// Manager represents information about a manager.
	// Exported type embedded the unexported field user.
	type Manager struct {
		Title string

		user
	}

	// Create a value of type Manager from the users package.
	// During construction, we are only able to initialize the exported field Title. We cannot
	// access the embedded type directly.
	u := users.Manager{
		Title: "Dev Manager",
	}

	// However, once we have the manager value, the exported fields from that unexported type are
	// accessible.
	u.Name = "Hoanh"
	u.ID = 101

	fmt.Printf("User: %#v\n", u)

*/
