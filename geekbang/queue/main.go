package main

import "fmt"

type MyQueue struct {
	input  []int
	output []int
}

func Constructor() MyQueue {
	return MyQueue{[]int{}, []int{}}
}

func (this *MyQueue) Push(x int) {
	this.input = append(this.input, x)
}

func (this *MyQueue) Pop() int {
	if this.Empty() {
		return 0
	}
	ret := this.Peek()
	this.output = this.output[:len(this.output)-1]
	return ret
}

func (this *MyQueue) Peek() int {
	if this.Empty() {
		return 0
	}
	if len(this.output) == 0 {
		for len(this.input) > 0 {
			this.output = append(this.output, this.input[len(this.input)-1])
			this.input = this.input[:len(this.input)-1]
		}
	}
	return this.output[len(this.output)-1]
}

func (this *MyQueue) Empty() bool {
	return len(this.input) == 0 && len(this.output) == 0
}

func main() {
	fmt.Println("vim-go")
}
