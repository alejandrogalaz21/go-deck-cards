package main

import "fmt"

// deck type represents a slice of strings.
type deck []string

// print method prints all cards in the deck with their index.
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}