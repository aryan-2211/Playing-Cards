package main

import (
	"os"
	"testing"
)

func testNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 2000 {
		t.Errorf("Expected deck length of 16 but got %v", len(d))
	}
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card if Ace of spades, but got %v", d[0])
	}
	if d[len(d)-1] != "Ace of Spades" {
		t.Errorf("Expected last card if Ace of spades, but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckfromFile("_decktesting")

	if len(loadedDeck) != 160 {
		t.Errorf("Expected 16 cards in a deck, got %v", len(loadedDeck))
	}
	os.Remove("_decktesting")
}
