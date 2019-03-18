package main

import "fmt"

func removeElement(nums []int, val int) int {

	elem := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			continue
		}
		nums[elem] = nums[i]
		elem++
	}

	return elem
}

func main() {
	nums := []int{3, 2, 2, 3}
	nums1 := []int{0, 1, 2, 2, 3, 0, 4, 2}
	fmt.Println(removeElement(nums, 3))
	fmt.Println(removeElement(nums1, 2))
}
