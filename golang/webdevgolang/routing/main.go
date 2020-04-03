package main

import (
	"fmt"
	"io"
	"log"
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

func formHandler(w http.ResponseWriter, req *http.Request) {
	v := req.FormValue("q")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `
	<form method="post">
	 <input type="text" name="q">
	 <input type="submit">
	</form>
	<br>`+v)
}

func rootHandler(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	log.Println(ctx)
	fmt.Printf("%+v\n", ctx)
}

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/dog/", dogHandler)
	http.HandleFunc("/cat", catHandler)
	http.HandleFunc("/search", handleURLSearchQuery)
	http.HandleFunc("/favicon.io", http.NotFound)
	http.HandleFunc("/form", formHandler)

	http.ListenAndServe(":8080", nil)
}
