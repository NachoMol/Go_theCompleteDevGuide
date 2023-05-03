package main

import "fmt"

// Create a new type of deck
// which is a slice of strings

type deck []string

func newDeck() deck {
	cards := deck{}
	cardsuits := []string{"Spades", "Diamonds", "Hearts", "Cubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardsuits {
		for _, value := range cardValues {
			cards = append(cards, suit+" of "+value)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
