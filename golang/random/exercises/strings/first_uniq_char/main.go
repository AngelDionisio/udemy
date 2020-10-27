package main

import (
	"fmt"
	"math"
)

func main() {
	str := "leetcode"
	fmt.Println(firstUniqueChar(str))
}

// Given a string, find the first non-repeating character in it and return its index. If it doesn't exist, return -1.
// keep track of the indexes of each character encountered, if the character alraedy exists in the map, invalidate as
// a duplicate by setting the index to -1
func firstUniqueChar(s string) int {
	// keep track of indexes of unique runes, if a rune is duplicated, give it a value of -1
	m := make(map[rune]int)
	for idx, v := range s {
		_, found := m[v]
		if found {
			m[v] = -1
			continue
		}

		m[v] = idx
	}

	first := math.MaxInt32
	for _, v := range m {
		if v > -1 && v < first {
			first = v
		}
	}

	if first == math.MaxInt32 {
		return -1
	}

	return first
}
