package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T){
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length 52, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected Ace of Spades in first position but got, %v", d[0])
	}
}

func TestWriteToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()
	d.writeToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected deck length 52, but got %v", len(d))
	}

	os.Remove("_decktesting")
}