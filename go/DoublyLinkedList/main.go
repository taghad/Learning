package main

import "fmt"

type node struct {
	value int
	next *node
	parent *node
}

func (e *node) addNewNode(head *node, value int) *node{
	if head == nil{
		head = new(node)
		head.value = value
		return head
	}
	next := head
	for next.next !=nil {
		next = next.next
	}
	next.next = new(node)
	next.next.value = value
	return head
}

func (e *node) deleteNode(head *node, value int) *node{

}

func (e *node) printList(head *node)  {
	for next := head; next !=nil;  next = next.next{
		fmt.Println("-> " , next.value)
	}

}

func (e *node) printListViceVersa(head *node)  {

}
func main()  {

}
