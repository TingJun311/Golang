package main

import (
	"fmt"
	"app/ds/linklist"
	"app/ds/Stack"
)

func main() {
	// Linked list
	firstNode := linklist.CreateNode("New node")
	fmt.Println(firstNode.GetLastNode())

	// Linked list stack
	var stack Stack.Stack
	stack.Push("Stack one")
	stack.PrintStack()
	fmt.Println("ok")

    num := 4513
    res := float64(num) / 100
    res2 := fmt.Sprintf("%.2f", res)
    fmt.Println(res2)
    
    var nums []int
    for i := 1; i <= 100000000; i++ {
        nums = append(nums, i)
    }
    fmt.Println(binarySearch(nums, -1))
    fmt.Println(linearSearch(nums, -1))
}

func binarySearch(array []int, num int) (int) {
    middleIndex := len(array) / 2
    loops := 0
    var tempArray []int
    tempArray = array

    for middleIndex > 0 {
        loops++
        if num < tempArray[middleIndex - 1] {
            tempArray = tempArray[:middleIndex]
        }
        if num > tempArray[middleIndex - 1] {
            tempArray = tempArray[middleIndex:]
        }
        if tempArray[middleIndex - 1] == num {
            fmt.Println("loops perform: ", loops)
            return tempArray[middleIndex - 1]
        }
        middleIndex = len(tempArray) / 2
    }
    fmt.Println("loops perform and found none: ", loops)
    return 0
}

func linearSearch(array []int, num int) (int) {
    loops := 0
    for _, i := range array {
        loops++
        if i == num {
            fmt.Println("loops perform liner search: ", loops)
            return i
        }
    }
    fmt.Println("loops perform liner search and found none: ", loops)
    return 0
}
