package main

import (
	"fmt"
	"os"
	"strings"
)

type deck []string

// newDeck creates and returns a new deck of cards
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

// print prints out all the cards in the deck
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// deal splits a deck into two decks, one with the specified hand size and the other with the remaining cards
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// toString converts the deck to a single string with cards separated by commas
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// saveToFile saves the deck to a file with the specified filename
func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

// newDeckFromFile reads a deck from a file with the specified filename and returns it as a deck
func newDeckFromFile(filename string) deck {
	bs, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return deck(s)
}
