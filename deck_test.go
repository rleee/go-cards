package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	/**
	 * test newDeck Length
	 */
	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, bug got %v", len(d))
	}

	/**
	 * test First card value
	 */
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected Ace of Spades for the first card, but got %v", d[0])
	}

	/**
	 * test Last card value
	 */
	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected Four of Clubs to be the last card, but got %v", d[len(d)-1])
	}

}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	const dN string = "_decktesting"
	os.Remove(dN)

	d := newDeck()
	d.saveToFile(dN)

	loadDeck := newDeckFromFile(dN)

	if len(loadDeck) != 16 {
		t.Errorf("Expected 16 cards in deck, but got %v", len(loadDeck))
	}

	os.Remove(dN)
}
