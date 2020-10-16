package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

/*
	create a new type, deck, to house the set
	of cards in a deck of cards.
	For sake of ease a card will be a strig of
	the value and suit it belongs to.
*/
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

/*
	create a receiver function on the deck.
	This means any variable of type deck in our
	program/package now has access to the print
	function.
*/
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}
