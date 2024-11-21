package main

import "fmt"

func main() {
	cards := []string{"Ace of Diamonds", newCard()}

	for i, card := range cards { 
		fmt.Println(card + " " + string(i))
	}
}

func newCard() string {
	card := "Ace of Diamonds"
	return card
}