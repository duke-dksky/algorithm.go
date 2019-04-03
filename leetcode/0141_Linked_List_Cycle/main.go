package main

import (
	"algorithm.go/utils"
	"fmt"
)

func hasCycle(head *utils.ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	fast := head
	slow := head
	for slow.Next != nil && fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false

}

func main() {
	fmt.Println("vim-go")
}
