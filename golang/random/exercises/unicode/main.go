package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func main() {
	characters := "🧭7a"
	// range over runes in a string
	// idenfity using the unicode package if a code point rune value IsDigit or IsLetter
	for _, v := range characters {
		if unicode.IsDigit(v) {
			intVal, _ := strconv.Atoi(string(v))
			fmt.Printf("looping through digit, binary value: %v, intVal: %v, stringVal: %v\n", v, intVal, string(v))
			continue
		}

		if unicode.IsLetter(v) {
			fmt.Println("looping through string", v, string(v))
			continue
		}

		fmt.Printf("non digit or string character value: %v -> %v\n", v, string(v))
	}
}
