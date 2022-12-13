package main

import "fmt"

func main () {
	colors := map[string]string{
		"dog": "bark",
		"Cat": "Meow",
	}

	printMap(colors)
}

func printMap (d map[string]string) {
	for color, hexCode := range d {
		fmt.Println(color)
		fmt.Println(hexCode)
	}
}
