package main //executable type package

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()

}
