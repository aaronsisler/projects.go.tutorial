package main

import "fmt"

func main() {
	// Two forms of variable declaration
	// var card string = "Ace of Spades"
	// card := newCard("Ace of Spades")
	//  Reassignment of "card" to a new string
	// card = "Two of Spades"

	// fmt.Println(card)

	cards := newDeck()
	fmt.Println("cards")
	fmt.Println(cards)

	fmt.Println("Single cards")
	cards.print()
	cards.shuffle()

	// var firstHand deck
	// firstHand, cards = deal(cards, 3)

	// firstHand, cards := deal(cards, 3)

	fmt.Println("")
	// fmt.Println("First hand")
	// firstHand.print()
	// fmt.Println("")
	// fmt.Println("Remaining Cards")
	cards.print()

	// cards.saveToFile("my-card-file")

	// newCards := newDeckFromFile("my-card-file")
	// fmt.Println(newCards)
}
