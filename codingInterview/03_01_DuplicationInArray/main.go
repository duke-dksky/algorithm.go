package main

import "fmt"

func Duplication(nums []int) bool {
	if len(nums) == 0 {
		return false
	}

	for i := range nums {
		if nums[i] < 0 || nums[i] > len(nums)-1 {
			return false
		}
	}

	for i := 0; i < len(nums); i++ {
		for i != nums[i] {
			if nums[i] == nums[nums[i]] {
				return true
			}
			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
		}
	}

	return false
}

func main() {
	fmt.Println(Duplication([]int{2, 1, 3, 1, 4}))
	fmt.Println(Duplication([]int{2, 4, 3, 1, 4}))
	fmt.Println(Duplication([]int{2, 1, 3, 0, 4}))
}
