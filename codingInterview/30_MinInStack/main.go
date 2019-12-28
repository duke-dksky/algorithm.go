package main

import "fmt"

type MyStack struct {
	data []int
	min  []int
}

func NewMyStack() MyStack {
	return MyStack{[]int{}, []int{}}
}

func (s *MyStack) Push(x int) {
	s.data = append(s.data, x)
	if len(s.min) == 0 || x < s.min[len(s.min)-1] {
		s.min = append(s.min, x)
	} else {
		s.min = append(s.min, s.min[len(s.min)-1])
	}
}

func (s *MyStack) Pop() int {
	if s.Empty() {
		return 0
	}
	ret := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	s.min = s.min[:len(s.min)-1]
	return ret
}

func (s *MyStack) Top() int {
	if s.Empty() {
		return 0
	}
	return s.data[len(s.data)-1]
}

func (s *MyStack) Empty() bool {
	return len(s.data) == 0
}

func (s *MyStack) Min() int {
	if s.Empty() {
		return 0
	}
	return s.min[len(s.min)-1]
}

func main() {
	s := NewMyStack()
	s.Push(5)
	s.Push(2)
	s.Push(3)
	s.Push(1)
	fmt.Println(s.Min())
	s.Pop()
	s.Pop()
	s.Pop()
	fmt.Println(s.Min())
}
