package main

import "fmt"

func MinNumberInRotatedArray(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	left, right := 0, len(nums)-1
	for left < right {
		if nums[left] < nums[right] {
			return nums[left]
		}
		if nums[left] == nums[right] {
			return MinInOrder(nums, left, right)
		}

		mid := left + (right-left)>>1
		if nums[left] <= nums[mid] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return nums[left]

}

func MinInOrder(nums []int, left, right int) int {
	res := nums[left]
	for i := left + 1; i < right; i++ {
		if res > nums[i] {
			res = nums[i]
		}
	}
	return res
}

func main() {
	fmt.Println(MinNumberInRotatedArray([]int{2, 1}))
}
