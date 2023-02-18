package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// print valiable
func (d deck) print() []error {
	var errors []error
	for i, card := range d {
		_, err := fmt.Println(i, card)
		if err != nil {
			errors = append(errors, err)
		}
	}

	return errors
}

// return multiple slice
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// convert dect to string
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// convert string to byte
func (d deck) toBytes() []byte {
	return []byte(d.toString())
}

// save to file
func (d deck) saveToFile(filepath string) error {
	return ioutil.WriteFile(filepath, d.toBytes(), 0666)
}

// load deck values from filepath
func newDeckFromFile(filepath string) deck {

	bs, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	return strings.Split(string(bs), ",")
}

// suffle cards in deck
func (d deck) suffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i, _ := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
