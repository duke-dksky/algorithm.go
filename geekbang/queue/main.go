package main

import "fmt"

// two stack
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

// circular queue
type CircularQueue struct {
	items []int
	n     int
	head  int
	tail  int
}

func NewCircularQueue(n int) *CircularQueue {
	return &CircularQueue{
		items: []int{},
		n:     n,
	}
}

func (c *CircularQueue) Empty() bool {
	return c.head == c.tail
}

func (c *CircularQueue) Full() bool {
	return (c.tail+1)%c.n == c.head
}

func (c *CircularQueue) Enqueue(x int) {
	if c.Full() {
		return
	}
	c.items = append(c.items, x)
	c.tail = (c.tail + 1) % c.n
}

func (c *CircularQueue) Dequeue() int {
	if c.Empty() {
		return 0
	}
	ret := c.items[c.head]
	c.head = (c.head + 1) % c.n
	return ret
}

func main() {
	fmt.Println("vim-go")
}
