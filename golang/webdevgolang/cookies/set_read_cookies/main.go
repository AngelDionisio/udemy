package main

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
)

func main() {
	http.HandleFunc("/", set)
	http.HandleFunc("/read", read)
	http.HandleFunc("/abundance", abundance)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.ListenAndServe(":8080", nil)
}

func set(w http.ResponseWriter, req *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "my-cookie",
		Value: "these are the cookie contents",
	})

	fmt.Println("typeof:", reflect.TypeOf(w))
	fmt.Println(w, "COOKIE WRITTEN - CHECK YOUR BROWSER")
}

// abundance writes multiple cookies
func abundance(w http.ResponseWriter, req *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "general",
		Value: "cookie contents about general things",
	})

	http.SetCookie(w, &http.Cookie{
		Name:  "specific",
		Value: "cookie contents about specific things",
	})

	fmt.Println(w, "COOKIES WRITTEN - CHECK YOUR BROWSER")
}

func read(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("my-cookie")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		log.Println(err)
	} else {
		fmt.Fprintln(w, "COOKIE #1:", c)
	}

	c2, err := req.Cookie("general")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		log.Println(err)
	} else {
		fmt.Fprintln(w, "COOKIE #2:", c2)
	}

	c3, err := req.Cookie("specific")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		log.Println(err)
	} else {
		fmt.Fprintln(w, "COOKIE #3:", c3)
	}

}
