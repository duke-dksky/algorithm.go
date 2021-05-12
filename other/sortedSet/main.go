package main

import (
	"fmt"
	"math/rand"
)

type node struct {
	ch       [2]*node
	priority int
	val      int
}

func (o *node) cmp(b int) int {
	switch {
	case b < o.val:
		return 0
	case b > o.val:
		return 1
	default:
		return -1
	}
}

func (o *node) rotate(d int) *node {
	x := o.ch[d^1]
	o.ch[d^1] = x.ch[d]
	x.ch[d] = o
	return x
}

type treap struct {
	root *node
	cnt  int
}

func (t *treap) _insert(o *node, val int) *node {
	if o == nil {
		t.cnt++
		return &node{priority: rand.Int(), val: val}
	}
	if d := o.cmp(val); d >= 0 {
		o.ch[d] = t._insert(o.ch[d], val)
		if o.ch[d].priority > o.priority {
			o = o.rotate(d ^ 1)
		}
	}
	return o
}

func (t *treap) insert(val int) {
	t.root = t._insert(t.root, val)
}

func (t *treap) _erase(o *node, val int) *node {
	if o == nil {
		return nil
	}
	if d := o.cmp(val); d >= 0 {
		o.ch[d] = t._erase(o.ch[d], val)
		return o
	}
	t.cnt--
	if o.ch[1] == nil {
		return o.ch[0]
	}
	if o.ch[0] == nil {
		return o.ch[1]
	}
	d := 0
	if o.ch[0].priority > o.ch[1].priority {
		d = 1
	}
	o = o.rotate(d)
	o.ch[d] = t._erase(o.ch[d], val)
	return o
}

func (t *treap) erase(val int) {
	t.root = t._erase(t.root, val)
}

func (t *treap) lowerBound(val int) (lb *node) {
	for o := t.root; o != nil; {
		switch c := o.cmp(val); {
		case c == 0:
			lb = o
			o = o.ch[0]
		case c == 1:
			o = o.ch[1]
		case c == -1:
			return o
		}
	}
	return
}

func (t *treap) upperBound(val int) (ub *node) {
	for o := t.root; o != nil; {
		switch c := o.cmp(val); {
		case c == 0:
			o = o.ch[0]
		case c == 1:
			ub = o
			o = o.ch[1]
		case c == -1:
			return o
		}
	}
	return
}

func (t *treap) size() int {
	return t.cnt
}

func (t *treap) begin() int {
	ans := 0
	for o := t.root; o != nil; {
		ans = o.val
		o = o.ch[0]
	}
	return ans
}

func (t *treap) end() int {
	ans := 0
	for o := t.root; o != nil; {
		ans = o.val
		o = o.ch[1]
	}
	return ans
}

func main() {
	t := &treap{}
	t.insert(3)
	t.insert(6)
	t.insert(4)
	t.insert(4)
	fmt.Println(t.size())
	t.erase(4)
	fmt.Println(t.size())
}
