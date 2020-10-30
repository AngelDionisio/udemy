package main

import (
	"regexp"
	"strings"
	"unicode"
)

// IsPalindrome returs true if a string is a palindrome
func IsPalindrome(s string) bool {
	if s == "" {
		return true
	}

	xr := []rune(strings.ToLower(s))
	left := 0
	right := len(xr) - 1

	for left < right {
		l := xr[left]
		r := xr[right]

		if !isAlphanumeric(l) {
			left++
			continue
		}

		if !isAlphanumeric(r) {
			right--
			continue
		}

		if l != r {
			return false
		}

		left++
		right--

	}

	return true
}

// isAlphanumeric returns true if a rune is a letter or digit.
// including latin characters, e.g. 'ñ'.
func isAlphanumeric(r rune) bool {
	if unicode.IsLetter(r) || unicode.IsDigit(r) {
		return true
	}
	return false
}

// IsPalindromeTwo checks if a string is a palindrome
func IsPalindromeTwo(s string) bool {
	if len(s) == 0 {
		return false
	}
	if len(s) == 1 {
		return true
	}

	sanitized := sanitize(s)

	listOfRunes := []rune(sanitized)

	for i := 0; i < len(listOfRunes)/2; i++ {
		if listOfRunes[i] != listOfRunes[len(listOfRunes)-i-1] {
			return false
		}
	}

	return true
}

// sanitize helper function that removes special characters from strings, as well as lower casing the result
func sanitize(value string) string {
	reg, _ := regexp.Compile("[^A-Za-z0-9]+")
	safe := reg.ReplaceAllString(value, "")
	return strings.ToLower(strings.Trim(safe, ""))
}

func isAlphanumericNonSpecialChars(r rune) bool {
	if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') {
		return true
	}
	return false
}
