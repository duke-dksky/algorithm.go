package main

import "fmt"

type ListNode struct {
	val  int
	next *ListNode
}

func PrintListReversingly(head *ListNode) {
	if head == nil {
		return
	}
	for head.next != nil {
		PrintListReversingly(head.next)
	}
	fmt.Println(head.val)
}

func main() {
	fmt.Println("vim-go")
}
