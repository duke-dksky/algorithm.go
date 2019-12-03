package main

import "fmt"

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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func ZeroSubArray(n []int) int {
	if len(n) == 0 {
		return -1
	}
	sum := make([]int, len(n))
	for k, v := range n {
		if k == 0 {
			sum[0] = n[0]
		} else {
			sum[k] = sum[k-1] + v
		}
	}

	bubbleSort(sum)

	ret := sum[1] - sum[0]
	for k, _ := range sum {
		if k == 0 {
			continue
		}
		diff := sum[k] - sum[k-1]
		ret = min(diff, ret)
	}

	return ret

}

func main() {
	fmt.Println(ZeroSubArray([]int{1, -2, 3, 10, -4, 7, 2, 4, -5}))
}
