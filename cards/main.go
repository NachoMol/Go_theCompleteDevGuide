package main

func main() {

	// cards := newDeck()
	// cards.saveToFile("my_cards")
	// deckFromFile("my_cards")

	// hand, remainingCards := deal(cards, 5)

	// hand.print()
	// remainingCards.print()

	cards := deckFromFile("my_cards")
	cards.print()

}
