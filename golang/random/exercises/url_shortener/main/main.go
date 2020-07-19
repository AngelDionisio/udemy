package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/angeldionisio/udemy/golang/random/exercises/url_shortener"
)

func main() {
	mux := defaultMux()
	yamlFilePath := flag.String("yaml", "paths.yaml", "a yaml file with file format of list 'path, url'")
	flag.Parse()

	ymlBytes, err := urlshort.ReadYamlFile(*yamlFilePath)
	if err != nil {
		fmt.Println("error trying to ready yaml file" + err.Error())
	}

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	// Inline yaml file representation
	// 	yaml := `
	// - path: /urlshort
	//   url: https://github.com/gophercises/urlshort
	// - path: /urlshort-final
	//   url: https://github.com/gophercises/urlshort/tree/solution
	// `
	yamlHandler, err := urlshort.YAMLHandler(ymlBytes, mapHandler)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
