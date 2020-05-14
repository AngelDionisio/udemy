package main

import (
	"fmt"
	"regexp"
	"strings"
)

// IsPalindrome checks if a string is a palindrome
func IsPalindrome(s string) bool {
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

func main() {
	fmt.Println(IsPalindrome("Amore, Roma"))
}
