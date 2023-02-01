package main

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()
	//cards.saveToFile("my_cards")
	//hands, remainingCards := deal(cards, 5)
	//hands.print()
	//remainingCards.print()
	//fmt.Println(cards.toString())
}
