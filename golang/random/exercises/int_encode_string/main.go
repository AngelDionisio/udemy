package main

import (
	"fmt"
	"strconv"
	"strings"
)

// could you think of cases where the size of the output
// would be bigger than the input
// would you be able to write the corresponding decode function to result string
// how would the API definition for these two functions would look like
// how would the process of getting this code to production look like

func main() {
	s := "xxxxyyyzz"

	encodedResult := encode(s)
	fmt.Printf("encoded: %v\n", encodedResult)

	decodedResult := decode(encodedResult)
	fmt.Printf("decoded: %v\n", decodedResult)

}

// encode converts a string with a sequence of repeated characters into a
// smaller string.
// xxxxyyyzz => 4x3y2z
func encode(s string) string {
	m := countMap(s)

	resultString := ""
	for k, v := range m {
		s := fmt.Sprintf("%v%v", v, k)
		resultString = resultString + s
	}

	return resultString
}

// decodes converts an encoded string back to its original form
func decode(s string) string {
	result := ""
	var lastNum int64
	for i := 0; i < len(s); i++ {
		if n, err := strconv.ParseInt(string(s[i]), 10, 64); err == nil {
			lastNum = n
			continue
		}

		result += strings.Repeat(string(s[i]), int(lastNum))
	}
	return result
}

// counts the number of occurences of each rune in a string
func countMap(s string) map[string]int {
	count := make(map[string]int)

	// iterate over runes in string
	for _, word := range s {
		count[string(word)]++
	}

	return count
}
