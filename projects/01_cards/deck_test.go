package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	expectedDeckLength := 16
	if len(d) != expectedDeckLength {
		t.Errorf("Expected: %v for deck length but got: %v", expectedDeckLength, len(d))
	}

	expectedCardName := "Ace of Spades"
	if d[0] != expectedCardName {
		t.Errorf("Expected: %v, got: %v", expectedCardName, d[0])
	}

	expectedCardName = "Four of Clubs"
	if d[len(d)-1] != expectedCardName {
		t.Errorf("Expected: %v, got: %v", expectedCardName, d[0])
	}

}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	filename := "_decktesting.txt"
	os.Remove(filename)

	d := newDeck()
	d.saveToFile(filename)

	loadedDeck := newDeckFromFile(filename)

	expectedDeckLength := 16
	if len(loadedDeck) != expectedDeckLength {
		t.Errorf("Expected: %v for deck length loaded from file but got: %v", expectedDeckLength, len(d))
	}

	os.Remove(filename)
}
