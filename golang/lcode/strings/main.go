package main

import (
	"fmt"

	"github.com/angeldionisio/udemy/golang/lcode/strings/palindrome"
)

func main() {
	str1 := "A man, a plan, a canal: Panama"
	isValid := palindrome.IsPalindrome(str1)
	fmt.Println(isValid)

	for _, v := range "0123456789" {
		fmt.Println(string(v), v)
	}
}
