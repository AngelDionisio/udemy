package main

import (
	"io"
	"net/http"
)

func dogHandler(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "dog dog dog")
}

func catHandler(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "cat cat cat")
}

func main() {
	http.HandleFunc("/dog/", dogHandler)
	http.HandleFunc("/cat", catHandler)

	http.ListenAndServe(":8080", nil)
}
