package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	str := "aaaaaaaaaaaaaaaaaaaaaaaaaa;aa##@@!!!ab133z;"
	fmt.Println("original:", str)
	encoded := encode(str)
	fmt.Println("encoded:", encoded)
	decoded := decode(encoded)
	fmt.Println("decoded:", decoded)
	fmt.Println(str == decode(encode(str)))
}

const delimiter = ';'

// encode :: string -> string
// aaaaaaaaaab133z
// 10;a;1;b;1;3;2;3;1;z
// keep a pointer on the current character, as long as next item in the loop is equal to the anchor
// increment counter, when a new character is found, set anchor to next index, restart counter to 1
// convert to slice of runes, as strings might take more than one byte to represent special chars
func encode(s string) string {
	if len(s) < 1 {
		return s
	}

	// anchor keeps track of a character, and its used to check if the next item repeats the char
	// at this anchor. Count keeps the amount of times the character has been repeated
	var sb strings.Builder
	var anchor, count int
	xsRunes := []rune(s)
	for idx, r := range s {
		if r == xsRunes[anchor] {
			count++
			continue
		}
		countString := strconv.Itoa(count)
		sb.WriteString(countString)
		sb.WriteRune(delimiter)
		sb.WriteRune(xsRunes[anchor])
		sb.WriteRune(delimiter)
		count = 1
		anchor = idx
	}
	countString := strconv.Itoa(count)
	sb.WriteString(countString)
	sb.WriteRune(delimiter)
	sb.WriteRune(xsRunes[anchor])

	return sb.String()
}

func decode(s string) string {
	var sb strings.Builder
	split := strings.Split(s, string(delimiter))
	fmt.Println(split)
	for i := 0; i < len(split); i = i + 2 {
		count, _ := strconv.Atoi(split[i])
		characterToPrint := split[i+1]

		s := strings.Repeat(characterToPrint, count)
		sb.WriteString(s)
	}
	return sb.String()
}
