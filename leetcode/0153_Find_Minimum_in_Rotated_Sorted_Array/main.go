package main

import "fmt"

func findMin(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}

	i, j := 0, len(nums)-1
	for i < j {
		mid := i + j>>2
		if nums[mid] 
	}

}

func main() {
	nums1 := []int{3, 4, 5, 1, 2}
	nums2 := []int{4, 5, 6, 7, 0, 1, 2}
	fmt.Println(findMin(nums1))
	fmt.Println(findMin(nums2))
}
