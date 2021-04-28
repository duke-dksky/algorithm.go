package main

import (
	"container/heap"
	"fmt"
)

type Item struct {
	val      int
	priority int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Pop() interface{} {
	res := (*pq)[len(*pq)-1]
	*pq = (*pq)[:len(*pq)-1]
	return res
}

func (pq *PriorityQueue) pop() *Item {
	return heap.Pop(pq).(*Item)
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*Item))
}

func (pq *PriorityQueue) push(x *Item) {
	heap.Push(pq, x)
}

func main() {
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	pq.push(&Item{
		val:      2,
		priority: 2,
	})
	pq.push(&Item{
		val:      1,
		priority: 1,
	})
	fmt.Println(pq.pop().val)
	fmt.Println(pq.pop().val)
}
