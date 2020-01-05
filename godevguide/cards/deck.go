package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Deck type to hold a slice of cards, so we can extend type with custom functionality
type Deck []string

// NewDeck creats a new suit of cards
// not adding a receiver here as when you are calling for a new deck there is no deck referece to chain to
func NewDeck() Deck {
	cards := Deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// Print logs to console the suit of cards
func (d Deck) Print() {
	for i, card := range d {
		fmt.Printf("%d: %s\n", i, card)
	}
}

// Deal cards
func Deal(d Deck, handSize int) (Deck, Deck) {
	return d[:handSize], d[handSize:]
}

// ToString flattens a Deck into a single string
func (d Deck) ToString() string {
	return strings.Join([]string(d), ",")
}

// SaveToFile writes deck to disk
func (d Deck) SaveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.ToString()), 0666)
}

// NewDeckFromFile creates a deck from a file
func NewDeckFromFile(filename string) Deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")

	// type convert []string -> Deck, which is also a []string
	return Deck(s)
}

// Shuffle randomly changes order of deck
func (d Deck) Shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
