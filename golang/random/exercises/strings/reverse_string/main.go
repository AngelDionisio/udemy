package main

import "fmt"

func main() {
	str := "this is a test 😊"
	fmt.Println(reverse(str))
}

func reverse(s string) string {
	// split string into runes
	r := []rune(s)
	head, tail := 0, len(r)-1

	for head < tail {
		r[head], r[tail] = r[tail], r[head]
		head++
		tail--
	}
	return string(r)
}
