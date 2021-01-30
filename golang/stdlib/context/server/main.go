package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("handler started")
	defer log.Println("handler ended")

	time.Sleep(time.Second * 5)
	fmt.Fprintln(w, "hello")
}
