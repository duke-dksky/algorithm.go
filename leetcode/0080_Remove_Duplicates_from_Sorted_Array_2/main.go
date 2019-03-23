package main

import (
	"fmt"
)

func removeDuplicates(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}

	j, count := 0, 1
	for i := 1; i < len(nums); i++ {
		if nums[j] == nums[i] {
			count++
			if count > 2 {
				continue
			}
		} else {
			count = 1
		}
		j++
		nums[j] = nums[i]
	}

	return j + 1
}

func main() {
	nums1 := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	nums2 := []int{1, 1, 1, 2, 2, 3}
	nums3 := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(removeDuplicates(nums1))
	fmt.Println(removeDuplicates(nums2))
	fmt.Println(removeDuplicates(nums3))
}
