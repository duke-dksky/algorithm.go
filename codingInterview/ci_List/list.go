package main

import (
	"fmt"
)

type ListNode struct {
	Value int
	Next  *ListNode
}

func NewListNode(value int) *ListNode {
	return &ListNode{value, nil}
}

func AddToTail(head *ListNode, value int) *ListNode {
	listNode := NewListNode(value)
	if head == nil {
		head = listNode
	} else {
		newHead := head
		for newHead.Next != nil {
			newHead = newHead.Next
		}
		newHead.Next = listNode
	}
	return head
}

func RemoveListNode(head *ListNode, value int) *ListNode {
	if head == nil {
		return nil
	}
	if head.Value == value {
		return head.Next
	}
	newHead := head
	for newHead.Next != nil && newHead.Next.Value != value {
		newHead = newHead.Next
	}
	if newHead.Next != nil && newHead.Next.Value == value {
		newHead.Next = newHead.Next.Next
	}
	return head
}

func PrintList(head *ListNode) {
	for head != nil {
		fmt.Println(head.Value)
		head = head.Next
	}
}
