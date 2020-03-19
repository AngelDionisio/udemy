package main

import (
	"fmt"
)

func main() {
	// to run 'clear; go run .\unique.go .\main.gogo run .\unique.go .\main.go'
	cases := []string{"angel", "dupess"}

	for _, v := range cases {
		fmt.Println(v, "contains unique runes?", Unique(v))
	}

}
