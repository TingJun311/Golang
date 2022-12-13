package main

import (
	"os"
	"testing"
)
func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 4 {
		t.Errorf("Expected deck of length 20 now %v", len(d))
	}
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first element of the deck its Ace of spades")
	}

	if d[len(d) - 1] != "Four of Clubs" {
		t.Errorf("Exected last card of the deck its Four of Clubs")
	}
}

func TestSaveToDeckAndRemoveTheDeck(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")
	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 but its not")
	}

	os.Remove("_decktesting")
}
