package main

import (
	"fmt"
	"strconv"
	"strings"
)

const separator = "_"

func main() {
	s := "xxxxyyywww2222"
	encoded := encodeString(s)
	fmt.Printf("%#v\n", encoded)

	decoded := decodeString(encoded)
	fmt.Printf("%v\n", decoded)
}

// xxxxyyyzzzzz -> 4x3y5z
func encodeString(s string) string {
	m := mapWordOccurences(s)
	fmt.Printf("map of occurrences: %#v\n", m)

	result := ""

	for k, v := range m {
		strToAppend := fmt.Sprintf("%v%v%v", v, k, separator)
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
	fmt.Println(listOfKeys)

	ret := ""
	for _, v := range listOfKeys {
		count, _ := strconv.Atoi(string(v[0]))
		char := string(v[1])
		toAppend := strings.Repeat(char, count)
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
