package main

import "fmt"

type MyStack struct {
	items []int
}

func Constructor() MyStack {
	return MyStack{[]int{}}
}

func (s *MyStack) Push(x int) {
	s.items = append(s.items, x)
}

func (s *MyStack) Pop() int {
	if s.Empty() {
		return 0
	}
	ret := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return ret
}

func (s *MyStack) Top() int {
	if s.Empty() {
		return 0
	}
	return s.items[len(s.items)-1]
}

func (s *MyStack) Empty() bool {
	return len(s.items) == 0
}

func main() {
	obj := Constructor()
	obj.Push(1)
	obj.Push(2)
	obj.Push(3)
	obj.Push(4)
	fmt.Println(obj.Pop())
	fmt.Println(obj.Top())
	fmt.Println(obj.Empty())
}
