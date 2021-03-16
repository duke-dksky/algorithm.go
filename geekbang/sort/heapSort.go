package main

import (
	"fmt"
)

// 从上往下进行调整
func minHeapAdjust(heap []int, i int, n int) {
	if heap == nil {
		return
	}
	j := i*2 + 1
	tmp := heap[i]

	// j < n 表示至少有左子节点
	for j < n {
		// 如果有右子节点，且值比左子节点小, 更新j
		if j+1 < n && heap[j+1] < heap[j] {
			j += 1
		}
		if heap[j] > tmp {
			break
		}
		// 将小值节点上移
		heap[i] = heap[j]
		i = j
		j = i*2 + 1
	}
	heap[i] = tmp

}

func heapSort(heap []int) {
	n := len(heap)
	if n < 2 {
		return
	}
	// 建堆
	for i := n/2 - 1; i >= 0; i-- {
		minHeapAdjust(heap, i, n)
	}
	// 堆排序
	for i := n - 1; i >= 1; i-- {
		heap[0], heap[i] = heap[i], heap[0]
		minHeapAdjust(heap, 0, i)
	}
}

func main() {
	a := []int{6, 5, 3, 1, 2, 7, 8}
	heapSort(a)
	fmt.Println(a)
}
