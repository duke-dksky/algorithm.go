package main

func countingSort(n []int) {
	if n == nil || len(n) <= 1 {
		return
	}

	// 最大值,决定桶的个数
	maxNum := n[0]
	for _, v := range n {
		if maxNum < v {
			maxNum = v
		}
	}

	// 在桶中,每个数值的个数
	bucket := make([]int, maxNum+1)
	for _, v := range n {
		bucket[v]++
	}

	// 在桶中为小于等于k的个数
	for k := 1; k <= maxNum; k++ {
		bucket[k] = bucket[k-1] + bucket[k]
	}

	// 排序, 为了稳定排序，从后向前遍历
	tmp := make([]int, len(n))
	for i := len(n) - 1; i >= 0; i-- {
		idx := bucket[n[i]] - 1
		tmp[idx] = n[i]
		bucket[n[i]]--
	}
	for i, v := range tmp {
		n[i] = v
	}
}
