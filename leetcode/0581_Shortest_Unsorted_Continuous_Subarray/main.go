package main

import (
	"fmt"
)

func findUnsortedSubarray(nums []int) int {
	if len(nums) == 0 || len(nums) == 1 {
		return 0
	}

	// 1.
	start, end := 0, len(nums)-1
	for start < end && nums[start] <= nums[start+1] {
		start++
	}
	for start < end && nums[end-1] <= nums[end] {
		end--
	}

	// 2. 在end和start之间的最大值和最小值
	minNum, maxNum := nums[start], nums[end]
	for i := start; i <= end; i++ {
		if nums[i] < minNum {
			minNum = nums[i]
		} else if nums[i] > maxNum {
			maxNum = nums[i]
		}
	}
	fmt.Println(minNum, maxNum)

	// 3.
	low, high := 0, len(nums)-1
	for low < start && nums[low] <= minNum {
		low++
	}
	for end < high && nums[high] >= maxNum {
		high--
	}

	if low == high {
		return 0
	}
	return high - low + 1

}

func main() {
	nums := []int{2, 6, 4, 8, 10, 9, 15}
	nums1 := []int{2, 3, 3, 2, 4}
	nums2 := []int{2, 3, 3, 3, 4}
	fmt.Println(findUnsortedSubarray(nums))
	fmt.Println(findUnsortedSubarray(nums1))
	fmt.Println(findUnsortedSubarray(nums2))
}
