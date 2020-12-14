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
	newDeck := deck{}

	cardSuites := []string{"Diamonds", "Hearts", "Clubs", "Spades"}
	cardValues := []string{"Ace", "Two", "Three", "Jack", "Queen", "King"}

	for _, suit := range cardSuites {
		for _, value := range cardValues {
			newDeck = append(newDeck, value+" of "+suit)
		}
	}

	return newDeck
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) toString() string {
	return strings.Join(d, ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	deckString := string(bs)                    // Ace of Spades,Two of Spades,Three of Spades...
	deckSlice := strings.Split(deckString, ",") // ["Ace of Spades" "Two of Spades"]

	return deckSlice
}

func (d deck) shuffle() {
	dLen := len(d)
	for i := range d {
		rand.Seed(time.Now().UnixNano())
		newPos := rand.Intn(dLen - 1)

		d[i], d[newPos] = d[newPos], d[i]
	}
}
