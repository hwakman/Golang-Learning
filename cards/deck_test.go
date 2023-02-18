package main

import (
	"os"
	"strings"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expect len to be 16 but got %d", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expect the first elm to be %s but got %s", "Ace of Spades", d[0])
	}

	if d[len(d) - 1] != "Four of Clubs" {
		t.Errorf("Expect the last elm to be %s but got %s", "Four of Clubs", d[len(d) - 1])
	}
}

func TestPrint(t *testing.T) {
	d := newDeck()

	errors := d.print()
	if len(errors) >= 1 {
		t.Errorf("Expect error to less than 1 got %d", len(errors))
	}
}

func TestDeal(t *testing.T) {
	d := newDeck()
	cards, remaining := deal(d, 3)
	if len(cards) > 3 {
		t.Errorf("Expect cards to less than 3 got %d", len(cards))
	}

	if len(remaining) <= 3 {
		t.Errorf("Expect remaining to gether than 3 got %d", len(remaining))
	}
}

func TestToString(t *testing.T) {
	d := newDeck()

	deckString := d.toString()
	if deckString != strings.Join(d, ",") {
		t.Errorf("Invalid value return %s", deckString)
	}
}

func TestToBytes(t *testing.T) {
	d := newDeck()
	deckByte := d.toBytes()
	if deckByte == nil {
		t.Errorf("Invalid value return %s", deckByte)
	}
}

func TestSaveToFileAndReadDeckFromFile(t *testing.T) {
	filename := "_decktesting"

	// cleanup 
	os.Remove(filename)

	d := newDeck()
	d.saveToFile(filename)

	loadDeck := newDeckFromFile(filename)
	if len(loadDeck) != 16 {
		t.Errorf("Expect len to be 16 but got %d", len(d))
	}

	if loadDeck[0] != "Ace of Spades" {
		t.Errorf("Expect the first elm to be %s but got %s", "Ace of Spades", loadDeck[0])
	}

	if loadDeck[len(d) - 1] != "Four of Clubs" {
		t.Errorf("Expect the last elm to be %s but got %s", "Four of Clubs", loadDeck[len(d) - 1])
	}

	os.Remove(filename)
}

func TestSuffle(t *testing.T) {
	d := newDeck()

	firstElement := d[0]
	lastElement := d[len(d) - 1]
	if firstElement != "Ace of Spades" {
		t.Errorf("Expect the first elm to be %s but got %s", "Ace of Spades", firstElement)
	}

	if lastElement != "Four of Clubs" {
		t.Errorf("Expect the last elm to be %s but got %s", "Four of Clubs", lastElement)
	}

	// suffle 
	d.suffle()
	if d[0] == firstElement {
		t.Errorf("Data should be suffle got first element %s", d[0])
	}

	if d[len(d) - 1] == lastElement {
		t.Errorf("Data should be suffle got last element %s", d[len(d) - 1])
	}
}
