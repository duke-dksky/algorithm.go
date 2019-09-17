package main

import "fmt"

func FindMinValInLoopArray(a []int) int {
	if a == nil || len(a) == 0 {
		return -1
	}
	left, right := 0, len(a)-1
	for left < right {
		mid := left + (right-left)/2
		if a[right] > a[mid] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return a[left]
}

func FindMinValInLoopArray1(a []int) int {
	if a == nil || len(a) == 0 {
		return -1
	}
	left, right := 0, len(a)-1
	for left < right {
		if a[left] < a[right] {
			return a[left]
		}
		mid := left + (right-left)/2
		if a[left] < a[mid] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return a[left]
}

func main() {
	fmt.Println(FindMinValInLoopArray1([]int{4, 5, 6, 7, 8, 0, 1, 2, 3}))
}
