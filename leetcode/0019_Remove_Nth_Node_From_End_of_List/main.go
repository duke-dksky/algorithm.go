package main

import (
	"algorithm.go/utils"
	//"fmt"
)

// 要修改原链表要传指针的指针,这个是通过返回值返回的
func removeNthFromEnd(head *utils.ListNode, n int) *utils.ListNode {

	fastHead, slowHead := head, head
	var slowPreHead *utils.ListNode
	for i := 0; i < n-1; i++ {
		fastHead = fastHead.Next
	}
	for fastHead.Next != nil {
		fastHead = fastHead.Next
		slowPreHead = slowHead
		slowHead = slowHead.Next
	}

	if slowPreHead != nil {
		slowPreHead.Next = slowHead.Next
	} else {
		return head.Next
	}

	return head
}

func main() {
	l := utils.NewListNode(1)
	utils.AddValueToListNode(&l, 2)
	utils.AddValueToListNode(&l, 3)
	utils.AddValueToListNode(&l, 4)
	utils.AddValueToListNode(&l, 5)
	utils.PrintListNode(removeNthFromEnd(l, 5))
}
