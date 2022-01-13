package main

func main() {

	cards := newDeck()

	cards.saveToFile("cards")

	cards.shuffle()

	cards.print()

	newDeckOfCards := loadFromFile("cards")

	newDeckOfCards.print()

	hand, remainingCards := deal(cards, 5)

	hand.print()

	remainingCards.print()
}
