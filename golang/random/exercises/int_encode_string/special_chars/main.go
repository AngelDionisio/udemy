package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	str := "aaaaaaaaaaaaaaaaaaaaaaaaaa;aa##@@!!!ab133zz;"

	encoded := encode(str)
	decoded := decode(encoded)

	fmt.Println("original:", str)
	fmt.Println("encoded:", encoded)
	fmt.Println("decoded: ", decoded)
	fmt.Println(str == decode(encode(str)))
}

const delimiter = ';'

// encode :: string -> string
// abaaaaaaaaaatz333322
func encode(s string) string {
	if len(s) == 0 || len(s) == 1 {
		return s
	}

	var sb strings.Builder
	xRunes := []rune(s)
	anchor, count := 0, 0
	stringTemplate := "%v%v;"
	for idx, r := range s {
		if r == xRunes[anchor] {
			count++
			continue
		}
		sb.WriteString(fmt.Sprintf(stringTemplate, string(xRunes[anchor]), count))
		count = 1
		anchor = idx
	}

	sb.WriteString(fmt.Sprintf(stringTemplate, string(xRunes[anchor]), count))

	return sb.String()
}

func decode(s string) string {
	if len(s) == 0 || len(s) == 1 {
		return s
	}

	var sb strings.Builder
	xRunes := []rune(s)
	currentRune := xRunes[0]
	for i := 1; i < len(xRunes); i++ {
		count := ""
		for i < len(xRunes) && xRunes[i] != delimiter {
			count += string(xRunes[i])
			i++
		}

		repeat, err := strconv.Atoi(count)
		if err != nil {
			fmt.Println("error due to:", err)
		}
		for repeat > 0 {
			sb.WriteRune(currentRune)
			repeat--
		}
		if i+1 >= len(xRunes) {
			break
		}
		if i < len(xRunes) {
			i++ // move away from comma
			currentRune = xRunes[i]
		}

	}

	return sb.String()
}
