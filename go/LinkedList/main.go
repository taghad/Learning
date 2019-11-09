package main

import "fmt"

type node struct {
	value int
	next *node
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
	if head.value == value {
		if  head.next == nil{
			return nil
		}
		head.value = head.next.value
		head.next = head.next.next
		return head
	}
	for next := head; next !=nil;  next = next.next {
			if next.next.value == value {
				next.next = next.next.next
				return head
			}
	}
	return head

}

func (e *node) printList(head *node)  {
	for next := head; next !=nil;  next = next.next{
		fmt.Println("-> " , next.value)
	}

}
func main()  {
	var head  *node
	head = head.addNewNode(head,12)
	head = head.addNewNode(head,11)
	head = head.addNewNode(head,10)
	head = head.addNewNode(head,9)
	head = head.addNewNode(head,8)
	head = head.addNewNode(head,7)
	head = head.deleteNode(head,12)
	head = head.deleteNode(head,9)
	head = head.deleteNode(head,7)

	head.printList(head)
}
