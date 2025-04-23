package main

import "fmt"

func main() {
	// var card string = "Ace of Spades"
	card := "Ace of Spades"
	card = newCard("Two")
	fmt.Println(card)

	cards := deck{newCard("Five"), newCard("One"), newCard("Three"), "April fool bnaya"}
	cards = append(cards, newCard("No cards"), "Testing")

	// iterating through the cards
	for index, eachCard := range cards {
		fmt.Println(index, eachCard)
	}

	// fmt.Println(cards)

}

func newCard(cnt string) string {
	return cnt + " of Daimonds"
}
