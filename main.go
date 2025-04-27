package main

import "fmt"

func main() {
	// Converting string to byte slice
	card := "Ace of Spades"
	fmt.Println([]byte(card))

	// creating a neew deck of cards
	cards := newDeck()
	cards = append(cards, "Testing")

	// Saving deck to file
	// cards.saveToFile("myDeck")

	// Reading deck from the file
	fmt.Println(newDeckFromFile("myDeck"))
}
