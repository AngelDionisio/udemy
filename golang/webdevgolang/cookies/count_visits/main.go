package main

import (
	"io"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", root)
	http.Handle("/favicon.io", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func root(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("my-cookie")

	if err == http.ErrNoCookie {
		cookie = &http.Cookie{
			Name:  "my-cookie",
			Value: "0",
		}
	}

	count, err := strconv.Atoi(cookie.Value)
	if err != nil {
		log.Fatalln(err)
	}
	count++
	cookie.Value = strconv.Itoa(count)

	http.SetCookie(w, cookie)

	io.WriteString(w, cookie.Value)
}
