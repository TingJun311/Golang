package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct {}
type spanishBot struct {}
type germanBot	struct {}
type italianBot struct {}

func main () {
	eb := englishBot {}
	sb := spanishBot {}
	gb := germanBot {}
	ib := italianBot {}

	printGreeting(eb)
	printGreeting(sb)
	printGreeting(gb)
	printGreeting(ib)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}


func (sb spanishBot) getGreeting () string {
	return "Hello there"
}

func (eb englishBot) getGreeting () string {
	return "Hola!"
}

func (gb germanBot) getGreeting () string {
	return "hallo"
}

func (ib italianBot) getGreeting () string {
	return "Italian 'Hello'"
}
