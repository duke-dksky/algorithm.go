package main

import "fmt"

func QuickSort(n []int) {
	if n == nil || len(n) <= 1 {
		return
	}
	quickSort(n, 0, len(n)-1)
}

func quickSort(n []int, begin, end int) {
	if begin >= end {
		return
	}
	idx := partition_2(n, begin, end)
	quickSort(n, begin, idx-1)
	quickSort(n, idx+1, end)
}

// 类似选择排序,分为`排序区`和`非排序区`
func partition(n []int, begin, end int) int {
	pivot := n[end]
	tmp := begin
	for i := begin; i <= end; i++ {
		if n[i] <= pivot {
			n[tmp], n[i] = n[i], n[tmp]
			tmp++
		}
	}
	return tmp - 1
}

func partition_2(n []int, begin, end int) int {
	pivot := n[begin]
	for begin < end {
		for begin < end && n[end] >= pivot {
			end--
		}
		n[begin] = n[end]
		for begin < end && n[begin] <= pivot {
			begin++
		}
		n[end] = n[begin]
	}
	n[begin] = pivot
	return begin
}

func main() {
	n := []int{6, 5, 4, 1, 2, 3}
	QuickSort(n)
	fmt.Println(n)
}
