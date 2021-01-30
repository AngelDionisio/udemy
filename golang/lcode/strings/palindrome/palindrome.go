package palindrome

import (
	"strings"
	"unicode"
)

// IsPalindrome determines if a string it is a palindrome,
// considering only alphanumeric characters and ignoring cases.
func IsPalindrome(s string) bool {
	s = strings.ToLower(s)
	xrunes := []rune(s)
	leftPtr, rightPtr := 0, len(xrunes)-1

	for leftPtr < rightPtr {
		l := xrunes[leftPtr]
		r := xrunes[rightPtr]
		if !isValidCharacter(l) {
			leftPtr++
			continue
		}

		if !isValidCharacter(r) {
			rightPtr--
			continue
		}

		if l != r {
			return false
		}

		leftPtr++
		rightPtr--
	}

	return true
}

func isValidCharacter(r rune) bool {
	if !unicode.IsLetter(r) || !unicode.IsDigit(r) {
		return false
	}

	return true
}
