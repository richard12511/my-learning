package main

func main() {
	cards := newDeckFromFile("cards")
	// hand, _ := deal(cards, 7)ß
	// cards.writeToFile("cards")ß

	cards.shuffle()
	cards.print()
}