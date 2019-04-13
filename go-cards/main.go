package main

func main() {

	cards := newDeck()
	cards.shuffle()
	cards.print()

	// cards := newDeckFromFile("my_card")
	// cards.print()

	// cards := newDeck()

	// cards.saveToFile("my_cards")

	// hand, remainingCards := deal(cards, 5)

	// hand.print()
	// remainingCards.print()

}
