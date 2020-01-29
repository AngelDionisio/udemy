package main

import "fmt"

func main() {
	cards := NewDeck()
	// cards.Shuffle()
	cards.Print()
	fmt.Println(len(cards))
}
