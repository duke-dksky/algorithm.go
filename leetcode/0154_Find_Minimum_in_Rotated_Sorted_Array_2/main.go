package main

import "fmt"

func findMin(nums []int) int {

	if nums == nil || len(nums) == 0 {
		return 0
	}

	start, end := 0, len(nums)-1

	for start < end {

		//结束条件
		if nums[start] < nums[end] {
			return nums[start]
		}

		mid := start + (end-start)>>2

		// 准确判断值在哪一边,并且确定是否包含mid的值
		if nums[mid] >= nums[start] {
			start = mid + 1
		} else {
			end = mid
		}

	}

	return nums[start]

}

func main() {
	nums1 := []int{3, 4, 5, 1, 2}
	nums2 := []int{4, 5, 0, 1, 2, 3}
	nums3 := []int{2, 1}
	nums4 := []int{4, 5, 6, 7, 0, 1, 2}
	nums5 := []int{2, 2, 2, 0, 0, 0, 0}
	fmt.Println(findMin(nums1))
	fmt.Println(findMin(nums2))
	fmt.Println(findMin(nums3))
	fmt.Println(findMin(nums4))
	fmt.Println(findMin(nums5))
}
