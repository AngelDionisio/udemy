package main

// Unique accepts a string and returns true if there are no
// repeated characters in the string, otherwise false is returned
func Unique(s string) bool {
	// convert string to []runes to properly compare characters
	xrune := []rune(s)
	// map to keep track of already seen runes
	seenMap := make(map[rune]bool)

	for _, v := range xrune {
		_, found := seenMap[v]
		if found {
			return false
		}

		seenMap[v] = true
	}

	return true
}
