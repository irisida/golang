package main

import "fmt"

func newCard() string {
	return "card value"
}

func main() {
	card := newCard()
	fmt.Println(card)
}
