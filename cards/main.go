
package main

import "fmt"

func main () {
	newCards := newDeck()
	newDeck, newDeck2 := deal(newCards, 2)
	newDeck.print()
	newDeck2.print()
	newbyte := []byte("Testing convert")
	for _, bytes := range newbyte {
		fmt.Println(bytes)
	}
	newCards.toString()
	fmt.Println(newCards)
	newCards.saveToFile("new_text.txt")
	newDeck3 := newDeckFromFile("new_text.txt")
	newDeck3.print()
	newDeck3.suffleDeck()
	newDeck3.print()
}

func newCard () string {
	return "This is a new string"
}
