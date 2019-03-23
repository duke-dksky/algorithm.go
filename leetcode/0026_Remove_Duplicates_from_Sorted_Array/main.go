package main

import "fmt"

func removeDuplicates(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}

	j := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[j] {
			continue
		}
		j++
		nums[j] = nums[i]
	}

	return j + 1
}

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(nums))
	fmt.Println(nums)
}
