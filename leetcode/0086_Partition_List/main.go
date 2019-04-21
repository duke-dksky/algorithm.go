package main

import (
	"algorithm.go/utils"
	"fmt"
)

func partition(head *utils.ListNode, x int) *utils.ListNode {

}

func main() {
	l := utils.NewListNode(1)
	utils.AddValueToListNode(&l, 4)
	utils.AddValueToListNode(&l, 3)
	utils.AddValueToListNode(&l, 2)
	utils.AddValueToListNode(&l, 5)
	utils.AddValueToListNode(&l, 2)
	utils.PrintListNode(l)
	utils.PrintListNode(partition(l, 3))
}
