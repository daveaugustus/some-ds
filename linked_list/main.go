package main

type Node struct {
	prev *Node
	next *Node
	key  interface{}
}

type List struct {
	head *Node
	tail *Node
}

func (l *List) Insert(key interface{}) {
	list := &Node{
		next: l.head,
		key:  key,
	}

	if l.head != nil {
		l.head.prev = list
	}
}

func main() {}
