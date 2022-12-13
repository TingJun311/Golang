package main

import "fmt"

func main () {
	color := map[string]string{
		"Red": "#ff000",
		"Green": "#ff123",
	}
	fmt.Println(color)
	setDiffMapValues(color, "Heredone")
	setDiffMapValues2(color, "Not done")
	fmt.Println(color)
}

func setDiffMapValues (c map[string]string, value string) {
	c["Red"] = value
}

func setDiffMapValues2 (c map[string]string, value string) {
	c["Green"] = value
}