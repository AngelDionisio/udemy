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

// http://localhost:8080/search?q=serve_me_this_data
func handleURLSearchQuery(res http.ResponseWriter, req *http.Request) {
	v := req.FormValue("q")
	io.WriteString(res, "Query value 'q':"+v)
}

func main() {
	http.HandleFunc("/dog/", dogHandler)
	http.HandleFunc("/cat", catHandler)
	http.HandleFunc("/search", handleURLSearchQuery)

	http.ListenAndServe(":8080", nil)
}
