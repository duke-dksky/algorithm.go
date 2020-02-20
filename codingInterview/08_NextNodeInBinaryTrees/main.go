package main

import "fmt"

type BinaryTreeNode struct {
	val    int
	left   *BinaryTreeNode
	right  *BinaryTreeNode
	parent *BinaryTreeNode
}

func GetNext(node *BinaryTreeNode) *BinaryTreeNode {
	if node == nil {
		return nil
	}

	if node.right != nil {
		rightNode := node.right
		for rightNode.left != nil {
			rightNode = rightNode.left
		}
		return rightNode
	}

	parentNode := node.parent
	for parentNode != nil && parentNode.right == node {
		node = parentNode
		parentNode = parentNode.parent
	}
	return parentNode
}

func main() {
	fmt.Println("vim-go")
}
