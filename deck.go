package main

import "fmt"

// deck type represents a slice of strings.
type deck []string

func newDeck() deck { 
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value + " of " + suit)
		}
	}

	return cards
}

// print method prints all cards in the deck with their index.
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i+1, card)
	}
}

