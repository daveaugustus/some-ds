package main

import (
	"fmt"
	"log"
)

type Node struct {
	value int
}

type Queue struct {
	Nodes []*Node
	count int
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Push(n *Node) {
	q.Nodes = append(q.Nodes, n)
	q.count++
}

func (q *Queue) Pop() {
	if q.count == 0 {
		log.Println("cannot remove a node from an empty queue")
	}
	q.count--
	q.Nodes = append(q.Nodes[0:0], q.Nodes[1:]...)
}

func main() {
	newQ := NewQueue()
	newQ.Push(&Node{1})
	newQ.Push(&Node{2})
	newQ.Push(&Node{3})
	newQ.Push(&Node{4})

	for i := range newQ.Nodes {
		fmt.Printf("%v\t", *&newQ.Nodes[i])
	}
	fmt.Println()

	newQ.Pop()
	newQ.Pop()
	newQ.Pop()

	for i := range newQ.Nodes {
		fmt.Printf("%v\t", *&newQ.Nodes[i])
	}
	fmt.Println()
}
