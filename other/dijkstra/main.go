package main

import (
	"container/heap"
	"math"
)

// 邻接表
// 1. 对于每个点,存储着一个列表,用来指向所有该点直接相连的点
// 2. 对于有权图,列表中元素的值对应这权重
type adjList struct {
	to int
	wt int
}

// edges[i] = [ui, vi, weighti]表示存在一条位于节点ui和vi之间的边,这条边的权重为weighti
func NewG(edges [][]int, n int) [][]adjList {
	g := make([][]adjList, n)
	for _, e := range edges {
		u, v, w := e[0], e[1], e[2]
		g[u] = append(g[u], adjList{v, w})
		g[v] = append(g[v], adjList{u, w})
	}
	return g
}

// 优先队列
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

// dijkstra
// 复杂度: (E+V)logV
// https://www.cnblogs.com/thousfeet/p/9229395.html
func dijkstra(g [][]adjList, start int) []int {
	dist := make([]int, len(g))
	for i := range dist {
		dist[i] = math.MaxInt64
	}
	dist[start] = 0
	// 优先队列
	pq := PriorityQueue{&Item{
		val:      start,
		priority: 0,
	}}
	for len(pq) > 0 {
		// 贪心策略: 在pq中选取一个离源点最近的节点弹出,即为确定节点
		item := pq.pop()
		v := item.val
		p := item.priority

		if dist[v] < p {
			continue
		}

		// 考察v的所有出边,做松弛操作
		for _, e := range g[v] {
			newD := dist[v] + e.wt
			if newD < dist[e.to] {
				dist[e.to] = newD
				pq.push(&Item{
					val:      e.to,
					priority: newD,
				})
			}
		}
	}
	return dist
}
