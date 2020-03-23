package main

import (
	"fmt"
	"net/http"
)

func getUser(req *http.Request) user {
	var u user

	// get cookie
	cookie, err := req.Cookie("session")
	if err != nil {
		return u
	}

	// if user exists already, get user
	if userName, ok := dbSessions[cookie.Value]; ok {
		u = dbUsers[userName]
	}

	fmt.Println("found user:", u)

	return u
}

func alreadyLoggedIn(req *http.Request) bool {
	cookie, err := req.Cookie("session")
	if err != nil {
		return false
	}

	userName := dbSessions[cookie.Value]
	_, ok := dbUsers[userName]
	return ok
}
