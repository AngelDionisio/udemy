package main

import (
	"fmt"
)

func main() {
	s1 := "nameless"
	s2 := "salesmen"
	res := IsAnagram("cinema", "iceman")
	res2 := IsAnagram(s1, s2)
	fmt.Printf("Is anagram? %v\n", res)
	fmt.Printf("Is anagram? %v\n", res2)
}

// IsAnagram accepts two strings and returns if one can spell target  using words from str
// a word, phrase, or name formed by rearranging the letters of another, such as cinema, formed from iceman.
// create mapping of target
// iterate through string, check against map, if the value is not in map, return false
// if it is, decrease the counter from map. Lastly check that all values in map are zero, if so, return true
func IsAnagram(s, target string) bool {
	if len(s) != len(target) {
		return false
	}

	tMap := mapOccurrences(target)
	for _, v := range s {
		// check if current rune is not in target map, or if count is < 0
		//  if so, return false, this means that there are more of said rune
		runeCount, found := tMap[v]
		if !found || runeCount < 0 {
			return false
		}
		tMap[v]--
	}

	// iterate through map, confirm all runes have been encountered
	for _, v := range tMap {
		if v != 0 {
			return false
		}
	}

	return true
}

func mapOccurrences(s string) map[rune]int {
	m := make(map[rune]int)
	for _, v := range s {
		m[v]++
	}
	return m
}
