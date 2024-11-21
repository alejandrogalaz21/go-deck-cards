package main

// main is the entry point for the program.
func main() {
	// Create a new deck of cards with an initial card.
	cards := deck{"Ace of Spaces"}
	// Add a new card to the deck.
	cards = append(cards, "Two of Hearts")
	// Print all cards in the deck.
	cards.print()
}

// newCard creates and returns a new card as a string.
func newCard() string {
	card := "Ace of Diamonds"
	return card
}