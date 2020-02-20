package main

import "fmt"

type ComplexListNode struct {
	value   int
	next    *ComplexListNode
	sibling *ComplexListNode
}

func CloneNodes(head *ComplexListNode) *ComplexListNode {
	tmpNode := head
	for tmpNode != nil {
		newNode := new(ComplexListNode)
		newNode.value = tmpNode.value
		newNode.next = tmpNode.next
		newNode.sibling = nil

		tmpNode.next = newNode

		tmpNode = newNode.next
	}
	return head
}

func ConnectSiblingNodes(head *ComplexListNode) {
	tmpNode := head
	for tmpNode != nil {
		cloneNode := tmpNode.next
		if tmpNode.sibling != nil {
			cloneNode.sibling = tmpNode.sibling.next
		}
		tmpNode = cloneNode.next
	}
}

func ReConnectNodes(head *ComplexListNode) *ComplexListNode {
	var cloneHead, cloneNode *ComplexListNode
	tmpNode := head

	if head != nil {
		cloneHead, cloneNode = head.next, head.next
		tmpNode.next = cloneNode.next
		tmpNode = tmpNode.next
	}

	for tmpNode != nil {
		cloneNode.next = tmpNode.next
		cloneNode = cloneNode.next

		tmpNode.next = cloneNode.next
		tmpNode = tmpNode.next
	}
	return cloneHead

}

func main() {
	fmt.Println("vim-go")
}
