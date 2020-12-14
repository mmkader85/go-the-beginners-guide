package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 24 {
		t.Errorf("NewDeck failed. Expected 24 but got %v", len(d))
	}

	if d[0] != "Ace of Diamonds" {
		t.Errorf("NewDeck failed. Expected 'Ace of Diamonds' but got '%v'", d[0])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	testFile := "_testdeck.txt"
	_ = os.Remove(testFile)

	d := newDeck()
	_ = d.saveToFile(testFile)

	_, err := os.Stat(testFile)
	if os.IsNotExist(err) {
		t.Errorf("Unable to save deck to file. %v", err)
	}

	if err == nil {
		nd := newDeckFromFile(testFile)
		if len(nd) != 24 {
			t.Errorf("NewDeck failed from File failed. Expected a deck but got %v", nd)
		}
	}

	_ = os.Remove(testFile)
}
