package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck lenght of 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected Age Of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected  Four of Clubs, but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	testDeckFileName := "_decktesting"
	os.Remove(testDeckFileName)

	deck := newDeck()
	deck.saveToFile(testDeckFileName)

	loadedDeck := newDeckFromFile(testDeckFileName)

	if len(loadedDeck) != 16 {
		t.Errorf("Expected deck lenght of 16, but got %v", len(loadedDeck))
	}

	os.Remove(testDeckFileName)
}
