package main

func main() {
	//deckFile := "RemainingCards.txt"

	card := newDeck()
	card.shuffle()
	card.print()

	//hand, remainingCards := deal(card, 5)
	//fmt.Println("Hand")
	//hand.print()
	//fmt.Println("Storing remaining cards in file")
	//err := remainingCards.saveToFile(deckFile)
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Println("Remaining cards saved successfully!")
	//}
	//
	//deckFromFile := newDeckFromFile(deckFile)
	//deckFromFile.print()
}