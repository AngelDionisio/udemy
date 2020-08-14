package main

import (
	"fmt"
	"strconv"
	"strings"
)

const separator = "_"

func main() {
	s := "xxxxyyywww2222222222"
	encoded := encodeString(s)
	fmt.Printf("encoded string: %#v\n", encoded)

	decoded := decodeString(encoded)
	fmt.Printf("decoded string: %v\n", decoded)
}

// xxxxyyyzzzzz -> 4x3y5z
func encodeString(s string) string {
	m := mapWordOccurences(s)
	fmt.Printf("map of occurrences: %#v\n", m)

	result := ""

	for k, v := range m {
		strToAppend := fmt.Sprintf("%v%v%v", k, v, separator)
		result += strToAppend
	}

	if len(result) > len(s) {
		fmt.Println("input string is smaller than result")
		return s
	}

	removedTrailingChar := result[:len(result)-1]
	return removedTrailingChar
}

func decodeString(s string) string {
	listOfKeys := strings.Split(s, separator)
	fmt.Println("list of keys:", listOfKeys)

	ret := ""
	for _, v := range listOfKeys {
		// position 0 will always correspond to the character we need to print
		// the rest, is the number of times the character appeared
		charToPrint := string(v[0])
		count, _ := strconv.Atoi(string(v[1:]))
		toAppend := strings.Repeat(charToPrint, count)
		ret += toAppend
	}
	return ret
}

func mapWordOccurences(s string) map[string]int {
	m := make(map[string]int)

	for _, v := range s {
		m[string(v)]++
	}

	return m
}
