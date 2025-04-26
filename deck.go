package main

import "fmt"

// create a new type of deck
// which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}
	cardTypes := []string{"Spades", "Clubs", "Diamonds", "Hearts"}
	cardNumber := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, eachCardcardType := range cardTypes {
		for _, eachCardNumber := range cardNumber {
			cards = append(cards, eachCardNumber+" of "+eachCardcardType)
		}
	}
	return cards
}

func (d deck) print() {
	for index, card := range d {
		fmt.Println(index, card)
	}
}

func deal(d deck, handSize int) []deck {
	// fmt.Println(d, handSize)
	decks := make([]deck, 2)
	decks[0] = d[:handSize]
	decks[1] = d[handSize:]

	return decks
}
