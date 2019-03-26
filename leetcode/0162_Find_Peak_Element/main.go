package main

import "fmt"

func findPeakElement(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return -1
	}

	start, end := 0, len(nums)-1

	for start <= end {
		mid := start + (end-start)>>2
		if mid > 0 && nums[mid] < nums[mid-1] {
			end = mid - 1
		} else if mid < len(nums)-1 && nums[mid] < nums[mid+1] {
			start = mid + 1
		} else {
			return mid
		}
	}

	return -1

}

func main() {
	nums1 := []int{1, 2, 3, 1}
	nums2 := []int{1, 2, 1, 3, 5, 6, 4}
	fmt.Println(findPeakElement(nums1))
	fmt.Println(findPeakElement(nums2))
}
