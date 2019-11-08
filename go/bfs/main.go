package main

import (
	"container/list"
	"fmt"
)

type Node struct {
	value  int
	child  []Node
	parent *Node
}

func main() {
	root := Node{
		value:  10,
		child:  make([]Node, 0),
		parent: nil,
	}
	key := 22

	root.child = append(root.child, Node{19, make([]Node, 0), &root})
	root.child = append(root.child, Node{22, make([]Node, 0), &root})

	queue := list.New()
	queue.PushBack(root)

	for queue.Len() > 0 {
		e := queue.Front().Value // First element
		if e.(Node).value == key {
			fmt.Println("we have this key in our tree")
			return
		}
		for i := 0; i < len(e.(Node).child); i++ {
			queue.PushBack(e.(Node).child[i])
		}
		queue.Remove(queue.Front()) // Dequeue
	}
	fmt.Println("we don't have this key in our tree")
}