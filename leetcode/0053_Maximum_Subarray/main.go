package main

import (
	"fmt"
)

func maxNum(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxSubArray(nums []int) int {

	curMax, finMax := 0, nums[0]
	for _, num := range nums {
		curMax += num
		finMax = maxNum(curMax, finMax)
		if curMax < 0 {
			curMax = 0
		}
	}
	return finMax

}

func main() {
	nums1 := []int{1, -2, 3, 10, -4, 7, 2, -5}
	nums2 := []int{-3, -2, -3, -10, -4, -7, -2, -5}
	nums3 := []int{-5, 4, 15, -10, -1, -2, 15, 16}
	nums4 := []int{1}
	fmt.Println(maxSubArray(nums1))
	fmt.Println(maxSubArray(nums2))
	fmt.Println(maxSubArray(nums3))
	fmt.Println(maxSubArray(nums4))
}
