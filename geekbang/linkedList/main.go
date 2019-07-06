package main

import "fmt"

type Object interface{}

type ListNode struct {
	Data Object
	Next *ListNode
}

func NewListNode(data Object) *ListNode {
	return &ListNode{data, nil}
}

func AddValueToList(head *ListNode, data Object) *ListNode {
	newListNode := NewListNode(data)
	if head == nil {
		head = newListNode
	} else {
		newHead := head
		for newHead.Next != nil {
			newHead = newHead.Next
		}
		newHead.Next = newListNode
	}
	return head
}

func PrintList(head *ListNode) {
	for head != nil {
		fmt.Println(head.Data)
		head = head.Next
	}
}

// leetcode 206
func ReverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var newHead, tmpHead *ListNode
	for head != nil {
		tmpHead = head.Next
		head.Next = newHead
		newHead = head
		head = tmpHead
	}
	return newHead
}

// leetcode 141
func IsExistLoop(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slowHead, fastHead := head, head
	for fastHead != nil && fastHead.Next != nil {
		slowHead = slowHead.Next
		fastHead = fastHead.Next.Next
		if slowHead == fastHead {
			return true
		}
	}
	return false
}

// leetcode 21
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	tmpHead := head
	for l1 != nil && l2 != nil {
		if l1.Data.(int) <= l2.Data.(int) {
			tmpHead.Next = l1
			l1 = l1.Next
		} else {
			tmpHead.Next = l2
			l2 = l2.Next
		}
		tmpHead = tmpHead.Next
	}
	if l1 == nil {
		tmpHead.Next = l2
	} else {
		tmpHead.Next = l1
	}

	return head.Next
}

// leetcode 19
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	fast, slow := head, head
	var slowPre *ListNode
	for i := 0; i < n-1; i++ {
		fast = fast.Next
	}
	for fast.Next != nil {
		fast = fast.Next
		slowPre = slow
		slow = slow.Next
	}
	if slowPre != nil {
		slowPre.Next = slow.Next
	} else {
		return head.Next
	}
	return head

}

// leetcode 876
func middleNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

func main() {
	var head *ListNode = nil
	head = AddValueToList(head, 1)
	head = AddValueToList(head, 2)
	PrintList(head)
	head = ReverseList(head)
	PrintList(head)
}
