package main

import (
	"fmt"
)

func main() {
	names := []string{"James Bond", "Moneypenny", "M"}

	upperCased := MapFunc(names, UpperCase)
	fmt.Printf("%#v\n", upperCased)
}
