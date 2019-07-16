package main

/*
1. 原地排序
2. 稳定排序
3. 时间复杂度：最好O(n),最坏O(n2),平均O(n2)
4. 交换次数为逆序度
*/
func bubbleSort(n []int) {
	if n == nil || len(n) <= 1 {
		return
	}
	for i := 0; i < len(n)-1; i++ {
		isSwap := false
		for idx := 0; idx < len(n)-1-i; idx++ {
			if n[idx] > n[idx+1] {
				n[idx], n[idx+1] = n[idx+1], n[idx]
				isSwap = true
			}
		}
		if !isSwap {
			break
		}
	}
}

/*
已排序区和未排序区
移动次数为逆序度

*/
func insertionSort(n []int) {
	if n == nil || len(n) <= 1 {
		return
	}
	for i := 1; i < len(n); i++ {
		tmpVal := n[i]
		idx := i - 1
		for ; idx >= 0; idx-- {
			if n[idx] > tmpVal {
				n[idx+1] = n[idx]
			} else {
				break
			}
		}
		// 包含idx=0的情况
		n[idx+1] = tmpVal
	}
}
