package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)


type deck []string
type laptopSize float64

func newDeck () deck {
	cards := deck {}
	cardSuits := []string {"Spades", "Diamonds", "Heart", "Clubs"}
	cardValue := []string {"Ace", "Two", "Three", "Four"}

	for _, card := range cardSuits {
		for _, value := range cardValue {
			cards = append(cards, value + " of " + card)
		}
	}
	return cards
}

func (d deck)  print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (v deck) toString() string {
	return strings.Join([]string(v), ",")
}

func (this laptopSize) getSizeOfLaptop() laptopSize {
	return this 
}

func deal(d deck, size int)  (deck, deck) {
	return d[:size], d[size:]
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile (filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	conDeck := strings.Split(string(bs), ",")
	return deck(conDeck)
}

func (d deck) suffleDeck () {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		randomNum := r.Intn(len(d) - 1)
		d[i], d[randomNum] = d[randomNum], d[i]
	}
}