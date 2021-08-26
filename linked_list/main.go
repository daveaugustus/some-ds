package main

import "fmt"

type Node struct {
	property int
	nextNode *Node
}

type LinkedList struct {
	headNode *Node
}

func (l *LinkedList) AddToHead(prop int) {
	node := Node{
		property: prop,
	}
	if node.nextNode == nil {
		node.nextNode = l.headNode
	}

	l.headNode = &node
}

func (ll *LinkedList) Display() {
	for i := ll.headNode; i != nil; i = i.nextNode {
		fmt.Println(i.property)
	}
}
func main() {
	ll := &LinkedList{}
	ll.AddToHead(1)
	ll.AddToHead(2)
	ll.AddToHead(3)
	ll.Display()
}
