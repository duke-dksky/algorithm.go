package main

import (
	"algorithm.go/utils"
	"fmt"
)

func detectCycle(head *utils.ListNode) *utils.ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	fast := head
	slow := head
	for slow.Next != nil && fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			slow = head
			for slow != fast {
				fast = fast.Next
				slow = slow.Next
			}
			return slow
		}
	}
	return nil
}

func main() {
	fmt.Println("vim-go")
}
