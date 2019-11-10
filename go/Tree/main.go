package main

import "fmt"

type node struct {
	value int
	left *node
	right *node
}

func insert(root **node, value int) {
	var nodeNow = *root
	if *root == nil {
		*root = new(node)
		(*root).value = value
		return
	}
	if value <= nodeNow.value {
		if nodeNow.left == nil {
			nodeNow.left = new(node)
			nodeNow.left.value = value
			return
		}
		insert(&nodeNow.left,value)
		return
	}
	if nodeNow.right == nil {
		nodeNow.right = new(node)
		nodeNow.right.value = value
		return
	}
	insert(&nodeNow.right,value)
}

func inOrderPrintTree(root *node)  {

	if root.left != nil {
		printTree(root.left)
	}
	{
		fmt.Println(root.value)
	}
	if root.right != nil {
		printTree(root.right)
	}

}

func main()  {


	var root *node

	insert(&root,10)
	insert(&root,5)
	insert(&root,3)
	insert(&root,6)
	insert(&root,13)
	insert(&root,1)

	inOrderPrintTree(root)



}