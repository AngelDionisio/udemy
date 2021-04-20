package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	s := "aaaaaaaaaa,,[]Abc1133"
	encoded := encode(s)
	decoded := decode(encoded)
	fmt.Printf("original: %v\n", s)
	fmt.Printf("encoded: %v\n", encoded)
	fmt.Printf("decoded: %v\n", decoded)
	same := s == decoded
	fmt.Println(same)
}

func encode(s string) string {
	if len(s) == 0 || len(s) == 1 {
		return s
	}

	var sb strings.Builder
	runes := []rune(s)
	anchor, count := 0, 0
	for idx, r := range s {
		if r == runes[anchor] {
			count++
			continue
		}

		sb.WriteRune(runes[anchor])
		sb.WriteRune(';')
		countString := strconv.Itoa(count)
		sb.WriteString(countString)
		sb.WriteRune(';')
		count = 1
		anchor = idx
	}
	sb.WriteRune(runes[anchor])
	sb.WriteRune(';')
	countString := strconv.Itoa(count)
	sb.WriteString(countString)

	return sb.String()
}

func decode(s string) string {
	separator := ";"
	if len(s) == 0 || len(s) == 1 {
		return s
	}

	list := strings.Split(s, separator)
	fmt.Println(list)

	var sb strings.Builder
	for i := 0; i < len(list); i += 2 {
		characterToPrint := list[i]
		repeatitions, _ := strconv.Atoi(list[i+1])

		for j := 0; j < repeatitions; j++ {
			sb.WriteString(characterToPrint)
		}
	}

	return sb.String()
}
