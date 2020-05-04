package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

// WordCount returns a map of word occurrences
func WordCount(s string) map[string]int {
	// clean string of special characters
	re, err := regexp.Compile(`[^\w]`)
	if err != nil {
		log.Fatal(err)
	}

	cleanStr := re.ReplaceAllString(s, " ")
	words := strings.Fields(cleanStr)
	resultMap := make(map[string]int)

	for _, word := range words {
		resultMap[word]++
	}

	return resultMap
}

func main() {
	str := "hello how are you doing?."
	wc := WordCount(str)

	fmt.Println(wc)

	// os.Stdout implements Writer
	w := os.Stdout

	fmt.Fprintf(w, "hello, writer\n")
}
