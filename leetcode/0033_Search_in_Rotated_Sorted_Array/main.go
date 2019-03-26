package main

import "fmt"

func search(nums []int, target int) int {

	if nums == nil || len(nums) == 0 {
		return -1
	}

	start, end := 0, len(nums)-1

	for start <= end {

		mid := start + (end-start)>>2
		if nums[mid] == target {
			return mid
		}

		// 准确判断值在哪一边,并且确定是否包含mid的值
		if nums[mid] >= nums[start] {
			if nums[start] <= target && target <= nums[mid] {
				end = mid
			} else {
				start = mid + 1
			}
		} else {
			if nums[mid] <= target && target <= nums[end] {
				start = mid
			} else {
				end = mid - 1
			}
		}

	}

	return -1

}

func main() {
	nums1 := []int{4, 5, 6, 7, 0, 1, 2}
	nums2 := []int{4, 5, 1, 2, 3}
	fmt.Println(search(nums1, 0))
	fmt.Println(search(nums2, 3))

}
