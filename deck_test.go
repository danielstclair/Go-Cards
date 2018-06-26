package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	dLength := len(d)
	if dLength != 24 {
		t.Errorf("Expected deck length of 24, but got %v", dLength)
	}
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}
	if d[dLength-1] != "Six of Clubs" {
		t.Errorf("Expected last card of Six of Clubs, but got %v", d[dLength-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_deckTesting")

	deck := newDeck()
	deck.saveToFile("_deckTesting")
	loadedDeck := newDeckFromFile("_deckTesting")

	if len(loadedDeck) != 24 {
		t.Errorf("Expected 16 cards in deck, got %v", len(loadedDeck))
	}

	os.Remove("_deckTesting")
}
