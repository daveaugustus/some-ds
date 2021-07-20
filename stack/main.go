package main

import (
	"fmt"
	"log"
)

type Node struct {
	value int
}

type Stack struct {
	nodes []*Node
	count int
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(n *Node) {
	s.nodes = append(s.nodes, n)
	s.count++
}

func (s *Stack) Pop() {
	if s.count == 0 {
		log.Println("can't remove node from an empty stack")
	}
	s.count--

	s.nodes = append(s.nodes[0:0], s.nodes[0:s.count]...)
}

func main() {
	newStack := NewStack()
	newStack.Push(&Node{4})
	newStack.Push(&Node{9})
	newStack.Push(&Node{5})
	newStack.Push(&Node{3})

	for i := range newStack.nodes {
		fmt.Printf("%v\t", *&newStack.nodes[i])
	}
	fmt.Println()

	newStack.Pop()
	newStack.Pop()

	for i := range newStack.nodes {
		fmt.Printf("%v\t", *&newStack.nodes[i])
	}
	fmt.Println()
}
