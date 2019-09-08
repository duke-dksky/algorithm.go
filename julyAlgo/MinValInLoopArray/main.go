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

func main() {
	fmt.Println(FindMinValInLoopArray([]int{4, 5, 6, 7, 0, 1, 2, 3}))
}
