package main

import "fmt"

func GetMissingNumber(data []int) int {
	if len(data) == 0 {
		return -1
	}
	left := 0
	right := len(data) - 1
	for left <= right {
		mid := left + (right-left)>>1
		if data[mid] != mid {
			if mid == 0 || (data[mid-1] == mid-1) {
				return mid
			}
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	if left == len(data) {
		return len(data)
	}
	return -1

}

func main() {
	fmt.Println(GetMissingNumber([]int{0}))
}
