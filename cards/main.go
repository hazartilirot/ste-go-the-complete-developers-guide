package main

func main() {

	cards := newDeck()

	saveToFile("cards", cards)

	newDeckOfCards := loadFromFile("cards")

	newDeckOfCards.print()

	hand, remainingCards := deal(cards, 5)

	hand.print()
	remainingCards.print()
}
