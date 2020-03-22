package main

import (
	"fmt"
	"github.com/satori/go.uuid"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", homeHandler)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func homeHandler(res http.ResponseWriter, req *http.Request) {
	// check if cookie is set, if not create with UUID
	cookie, err := req.Cookie("session-id")
	if err != nil {
		id, err := uuid.NewV4()
		if err != nil {
			log.Fatalf("could not create UUID: %v", err)
		}

		cookie = &http.Cookie{
			Name:  "session-id",
			Value: id.String(),
			// Secure:   true,
			HttpOnly: true,
		}

		http.SetCookie(res, cookie)
	}

	fmt.Println(cookie)
}
