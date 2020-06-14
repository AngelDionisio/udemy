package main

import (
	"fmt"
)

// user identifies an user in the system
type user struct {
	name  string
	email string
}

// notify implements a method for user. Sends notifications to user
// for different events
func (u *user) notify() {
	fmt.Printf("Sending user email to %q<%s>\n", u.name, u.email)
}

// admin is an user with superuser privileges
type admin struct {
	// person user // Not embedded
	user  // embedded type, all methods of user get promoted to the admin type.
	level string
}

func main() {
	ad := admin{
		user: user{
			name:  "James Bond",
			email: "bond@mi5.com",
		},
		level: "superadmin",
	}

	ad.user.notify()

	// because user is embedded in admin (inner type promotion)
	// admin can assess notify directly.
	ad.notify()
}
