package main

import (
	"fmt"
	"app/ds/linklist"
)

func main() {
	firstNode := linklist.CreateNode("New node")
	fmt.Println(firstNode.GetLastNode())
}
