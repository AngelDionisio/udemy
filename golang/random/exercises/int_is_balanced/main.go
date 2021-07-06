package main

import (
	"fmt"
)

func main() {
	s := "[({abc})]"
	isBal := IsBalanced(s)

	fmt.Println(isBal)
}

// IsBalanced iterates over each rune in a string. If the string is an open parenthesis
// it will store it in a stack, if it is a closed parenthesis, it will pop the last item
// in the slice, and compare them. If they are the same, it will set the list to a new list
// who has the last item removed.
func IsBalanced(s string) bool {
	if len(s) == 0 {
		return false
	}

	// slice to be used a stack to hold open parenthesis
	// upon encoutering a closed parens, will compare with the last open parens
	// in this list
	var list []rune

	for _, v := range s {
		if isOpenParens(v) {
			list = append(list, v)
		}

		if isClosingParens(v) {
			if len(list) < 1 {
				return false
			}
			lastItem := list[len(list)-1]
			if lastItem != matchOpen(v) {
				return false
			}
			list = list[:len(list)-1]
		}
	}

	return true
}

func matchOpen(s rune) rune {
	m := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	v := m[s]
	return v
}

func isOpenParens(s rune) bool {
	m := map[rune]bool{
		'(': true,
		'[': true,
		'{': true,
	}

	_, ok := m[s]
	if !ok {
		return false
	}

	return true
}

func isClosingParens(s rune) bool {
	m := map[rune]bool{
		')': true,
		']': true,
		'}': true,
	}

	_, ok := m[s]
	if !ok {
		return false
	}

	return true
}

func printSliceOfRunes(list []rune) {
	fmt.Println("Printing slice of rune contents:")
	for _, v := range list {
		fmt.Printf("rune value %v, string format: %v\n", v, string(v))
	}
}
