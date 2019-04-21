package utils

import (
	"fmt"
)

type Tree struct {
	Left  *Tree
	Data  int
	Right *Tree
}

func NewTree(v int) *Tree {
	return &Tree{nil, v, nil}
}

func InsertNodeToTree(t *Tree, v int) *Tree {
	if t == nil {
		return NewTree(v)
	}
	if v < t.Data {
		t.Left = InsertNodeToTree(t.Left, v)
		return t
	}
	t.Right = InsertNodeToTree(t.Right, v)
	return t
}

// reference: https://blog.csdn.net/monster_ii/article/details/82115772

// PreOrder
func PreOrder(t *Tree) {
	if t == nil {
		return
	}
	fmt.Printf("%d ", t.Data)
	PreOrder(t.Left)
	PreOrder(t.Right)
}

// 栈中元素都是自己和自己的左孩子都访问过了，而右孩子还没有访问到的节点
func PreOrderNonRecursive(t *Tree) {
	stack := make([]*Tree, 0)
	cur := t

	for cur != nil || len(stack) != 0 {
		for cur != nil {
			fmt.Printf("%d ", t.Data)
			stack = append(stack, cur)
			cur = cur.Left
		}
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		cur = top.Right
	}
}

// InOrder
func InOrder(t *Tree) {
	if t == nil {
		return
	}
	InOrder(t.Left)
	fmt.Printf("%d ", t.Data)
	InOrder(t.Right)
}

// 节点自身和它的右子树都没有被访问到的节点地址。
func InOrderNonRecursive(t *Tree) {
	stack := make([]*Tree, 0)
	cur := t

	for cur != nil || len(stack) != 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		fmt.Printf("%d ", top.Data)
		cur = top.Right
	}
}

// PastOrder
func PastOrder(t *Tree) {
	if t == nil {
		return
	}
	PastOrder(t.Left)
	PastOrder(t.Right)
	fmt.Printf("%d ", t.Data)
}

// 右子树和自身都没有被遍历到的节点,但是多一个last指针指向上一次访问到的节点，
// 用来确认是从根节点的左子树返回的还是从右子树返回的
func PastOrderNonRecursive(t *Tree) {
	stack := make([]*Tree, 0)
	cur := t
	var last *Tree

	for cur != nil || len(stack) != 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		top := stack[len(stack)-1]

		if top.Right == nil || top.Right == last {
			stack = stack[:len(stack)-1]
			fmt.Printf("%d ", top.Data)
			last = top
		} else {
			cur = top.Right
		}
	}
}

// levelOrder
func LevelOrder(t *Tree) {
	if t == nil {
		return
	}

	queue := make([]*Tree, 1)
	queue = append(queue, t)

	for len(queue) != 0 {
		top := queue[len(queue)-1]
		queue = queue[:len(queue)-1]

		if top.Left != nil {
			queue = append([]*Tree{top.Left}, queue...)
		}
		if top.Right != nil {
			queue = append([]*Tree{top.Right}, queue...)
		}

		fmt.Printf("%d ", top.Data)
	}
}
