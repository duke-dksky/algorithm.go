package utils

import (
	"fmt"
)

type ListNode struct {
	Data Object
	Next *ListNode
}

func NewListNode(data Object) *ListNode {
	return &ListNode{data, nil}
}

/* head必须为指针的指针,因为链表为空时,需要改变头指针 */
func AddValueToListNode(head **ListNode, value Object) {
	newListNode := NewListNode(value)
	if *head == nil {
		*head = newListNode
	} else {
		newHead := *head
		for newHead.Next != nil {
			newHead = newHead.Next
		}
		newHead.Next = newListNode
	}
	return
}

func RemoveNode(head **ListNode, value int) {
	if head == nil && *head == nil {
		return
	}

	if (*head).Data == value {
		*head = (*head).Next
	} else {
		newHead := *head
		for newHead.Next != nil && newHead.Next.Data != value {
			newHead = newHead.Next
		}
		if newHead.Next != nil && newHead.Next.Data == value {
			newHead.Next = newHead.Next.Next
		}
	}
	return
}

// print list node
func PrintListNode(head *ListNode) {
	if head == nil {
		return
	}
	newHead := head
	for newHead != nil {
		fmt.Printf("list node data: %v\n", newHead.Data)
		newHead = newHead.Next
	}
	return
}
