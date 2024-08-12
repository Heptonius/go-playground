package main

const FILE_NAME = "my_cards"

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()
}
