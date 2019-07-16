package main

func MergeSort(n []int) {
	if n == nil || len(n) <= 1 {
		return
	}
	mergeSort(n, 0, len(n)-1)
}

func mergeSort(n []int, begin, end int) {
	if begin >= end {
		return
	}
	mid := begin + (end-begin)/2
	mergeSort(n, begin, mid)
	mergeSort(n, mid+1, end)
	merge_2(n, begin, mid, end)
}

func merge(n []int, begin, mid, end int) {
	tmp := make([]int, end-begin+1)
	for i := begin; i <= end; i++ {
		tmp[i-begin] = n[i]
	}

	left, right := begin, mid+1

	for i := begin; i <= end; i++ {
		if left <= mid && right <= end {
			if tmp[left-begin] < tmp[right-begin] {
				n[i] = tmp[left-begin]
				left++
			} else {
				n[i] = tmp[right-begin]
				right++
			}
		} else if left <= mid {
			n[i] = tmp[left-begin]
			left++
		} else {
			n[i] = tmp[right-begin]
			right++
		}
	}

}

func merge_2(n []int, begin, mid, end int) {
	tmp := make([]int, end-begin+1)

	left, right := begin, mid+1

	for i := begin; i <= end; i++ {
		if left <= mid && right <= end {
			if n[left] < n[right] {
				tmp[i-begin] = n[left]
				left++
			} else {
				tmp[i-begin] = n[right]
				right++
			}
		} else if left <= mid {
			tmp[i-begin] = n[left]
			left++
		} else {
			tmp[i-begin] = n[right]
			right++
		}
	}

	for i, val := range tmp {
		n[begin+i] = val
	}
}
