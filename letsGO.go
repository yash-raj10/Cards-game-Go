package main

func main() {
	//cards := newDeck()

	//cards.printIt()

	//fmt.Print(cards.toString())

	//cards.saveToFile("yooo")

	//cards := newDeckFromFile("yooo")
	//cards.printIt()

	cards := newDeck()
	cards.shuffle()
	cards.printIt()
}
