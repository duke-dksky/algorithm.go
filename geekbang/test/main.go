package main

import (
	"fmt"
)

type MinStack struct {
	items    []int
	minItems []int
}

func Constructor() MinStack {
	return MinStack{[]int{}, []int{}}
}

func (s *MinStack) Push(x int) {
	s.items = append(s.items, x)
	if len(s.minItems) == 0 {
		s.minItems = append(s.minItems, x)
	} else {
		if s.minItems[len(s.minItems)-1] > x {
			s.minItems = append(s.minItems, x)
		} else {
			s.minItems = append(s.minItems, s.minItems[len(s.minItems)-1])
		}
	}
}

func (s *MinStack) Pop() int {
	if s.Empty() {
		return 0
	}
	s.minItems = s.minItems[:len(s.minItems)-1]
	ret := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return ret
}

func (s *MinStack) Top() int {
	if s.Empty() {
		return 0
	}
	return s.items[len(s.items)-1]
}

func (s *MinStack) Empty() bool {
	return len(s.items) == 0
}

func (s *MinStack) GetMin() int {
	if s.Empty() {
		return 0
	}
	return s.minItems[len(s.minItems)-1]
}

func main() {
	obj := Constructor()
	obj.Push(-2)
	obj.Push(0)
	obj.Push(-3)
	fmt.Println(obj.GetMin())
	obj.Pop()
	fmt.Println(obj.Top())
	fmt.Println(obj.GetMin())
}
