package main

import (
	"fmt"
	"net/http"
)

type myserver int

// ServeHTTP implements the Handler interface
func (s myserver) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "from serveHTTP")
}

func main() {
	var s myserver
	http.ListenAndServe(":8080", s)
}
