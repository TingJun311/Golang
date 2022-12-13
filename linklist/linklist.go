
package linklist

import (
	"fmt"
)

type Node struct {
	content interface{}
	next *Node
}


func CreateNode(content interface{}) (node Node) {
	var newNode Node

	newNode.content = content
	newNode.next = nil
	return newNode
}

func (n *Node) AddNode(node Node) {
	if n.next == nil {
		n.next = &node
	} else {
		tempPtr := n
		for tempPtr.next != nil {
			tempPtr = tempPtr.next
		}
		tempPtr.next = &node
	}
}

func (n *Node) GetLastNode() (node *Node){
	temp := n
	if n == nil {
		return
	}
	for temp.next != nil {
		temp = temp.next
	}
	return temp
}

func (n *Node) GetFirstNode() (node *Node) {
	temp := n
	firstNode := temp
	if n == nil {
		return
	}
	n = n.next
	return firstNode
}

func (n *Node) GetNodeByNum(num int) (content interface{}) {
	if n == nil {
		return
	}
	head := n
	numNodes := 1
	if num > head.Length() {
		return nil
	}
	for head.next != nil && numNodes < num {
		numNodes++
		head = head.next
	}
	return nil
}

func DeleteNode(nodePos int, headNode Node) (newNode *Node) {
	var prevNode *Node

	if nodePos == 0 {
		return
	}
	numNodes := 1
	cur := &headNode
	if nodePos == 1 {
		newNode := cur.next
		return newNode
	}
	for cur.next != nil && numNodes < nodePos {
		prevNode = cur
		cur = cur.next
		numNodes++
	}
	if cur.next != nil {
		prevNode.next = cur.next
	} else {
		prevNode.next = nil
	}
	newNode = &headNode
	return newNode
}


func (n *Node) SearchContent(content interface{}) (node Node, num int) {
	head := n
	if head == nil {
		return
	}
	num++
	for head.next != nil && head.content != content {
		num++
		head = head.next
	}
	if head.content == content {
		return CreateNode(head.content), num
	}
	return CreateNode(nil), 0
}

func InsertNode(headNode Node, index int, newNode Node) (*Node) {
	cur := &headNode
	temp := &headNode
	curIndex := 0
	if index == 0 {
		cur = &newNode
		cur.next = temp
		return cur
	}
	for cur.next != nil && curIndex < index - 1 {
		curIndex++
		cur = cur.next
	}
	if cur.next != nil {
		latestNode := cur.next
		cur.next = &newNode
		newNode.next = latestNode
	} else {
		cur.next = &newNode
	}
	cur = temp
	return cur
}

func (n *Node) Length() (int) {
	l := 1
	if n == nil {
		l = 0
		return l
	}
	temp := n
	for temp.next != nil {
		l++
		temp = temp.next
	}
	return l
}

func (n *Node) PrintList() {
	tempPtr := n
	for tempPtr.next != nil {
		fmt.Print(tempPtr.content, " -> ")
		tempPtr = tempPtr.next
	}
	fmt.Println(tempPtr.content)
}

//func main() {
//	linkList := createNode("first Node")
//	node2 := createNode("Second Node")
//	node3 := createNode("Third Node")
//	linkList.addNode(node2)
//	linkList.addNode(node3)
//	linkList.addNode(createNode("4 node"))
//	linkList.addNode(createNode("5 Node"))
//	linkList.addNode(createNode("6 node"))
//	newLinkList := deleteNode(2, linkList)
//	newLinkList.printList()
//	newLinkList2 := insertNode(linkList, 0, createNode("Beand New Node"))
//	newLinkList2.printList()
//}
