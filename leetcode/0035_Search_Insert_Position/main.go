package main

import "fmt"

func searchInsert(nums []int, target int) int {
	if nums == nil || len(nums) == 0 {
		return -1
	}

	i, j := 0, len(nums)-1

	for i <= j {
		mid := i + (j-i)>>2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			j = mid - 1
		} else {
			i = mid + 1
		}
	}
	return i

}

func main() {
	nums1 := []int{1, 3, 5, 6}
	fmt.Println(searchInsert(nums1, 5))
	fmt.Println(searchInsert(nums1, 2))
	fmt.Println(searchInsert(nums1, 7))
	fmt.Println(searchInsert(nums1, 0))
}
