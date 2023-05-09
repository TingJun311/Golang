package Stack

import (
	"fmt"
)

type Stack struct {
	content interface{}
	up *Stack
}

func CreateNode(content interface{}) (Stack){
	var newNode Stack
	if content != nil {
		newNode.content = content
		newNode.up = nil
	}
	return newNode
}

func (s *Stack) Push(content interface{}) {
	temp := s
	head := s
	newStackNode := CreateNode(content)
	if temp.content == nil {
		temp.content = content
		return
	}
	for temp.up != nil {
		temp = temp.up
	}
	temp.up = &newStackNode
	s = head
}

func InitEmptyStack() (*Stack) {
	var newStack Stack
	return &newStack
}

func (s *Stack) Pop() (content interface{}) {
	var prevNode *Stack
	length := 1

	head := s
	if head.content == nil {
		return nil
	} else {
		prevNode = head
	}
	for head.up != nil {
		prevNode = head
		head = head.up
		length++
	}
	getNodeContent := head.content
	if prevNode != nil {
		prevNode.up = nil
	} else if length == 1 {
		head.content = nil
	}
	return getNodeContent
}

func (s *Stack) CountStack() (int) {
	var count int

	temp := s
	for temp.up != nil {
		count++
		temp = temp.up
	}
	count++
	return count
}

func (s *Stack) PrintStack() {
	var tempSlice []interface{}
	temp := s
	for temp.up != nil {
		tempSlice = append(tempSlice, temp.content)
		temp = temp.up
	}
	tempSlice = append(tempSlice, temp.content)

	// Reverse a string
	for i, j := 0, len(tempSlice) - 1; i < j; i, j = i + 1, j - 1 {
		tempSlice[i], tempSlice[j] = tempSlice[j], tempSlice[i]
	}
	fmt.Printf("   ^   \n")
	for _, i := range tempSlice {
		fmt.Printf("   |   \n")
		fmt.Printf("%s \n", i)
	}
	fmt.Println("\n\n")
}

//func main() {
//	stack := InitEmptyStack()
//	stack.Push("testing")
//	stack.Push("testing 2")
//	stack.Push("testing 3")
//	stack.Push("testing 4")
//	stack.PrintStack()
//	fmt.Println(stack.CountStack())
//}
