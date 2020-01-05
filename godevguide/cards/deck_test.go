package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := NewDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %d", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be Ace of Spaces, but got %s", d[0])
	}

	if (d[len(d)-1]) != "King of Clubs" {
		t.Errorf("Expected first card to be King of Clubs, but got %s", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckTestFromFile(t *testing.T) {
	// if for some reason the file was not deleted on any iteration,
	// attempt to delete before test
	os.Remove("_decktesting")

	deck := NewDeck()
	deck.SaveToFile("_decktesting")

	loadedDeck := NewDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected 52 cards in deck, got %d", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
