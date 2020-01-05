package main

import "fmt"

func VerifySquenceOfBST(nums []int) bool {
	if len(nums) == 0 {
		return false
	}
	return VerifySquenceOfBSTRecursive(nums, 0, len(nums)-1)
}

func VerifySquenceOfBSTRecursive(nums []int, start, end int) bool {

	if start == end {
		return true
	}

	root := nums[end]
	i := start
	for ; i < end-1; i++ {
		if nums[i] > root {
			break
		}
	}
	for j := i; j < end-1; j++ {
		if nums[j] < root {
			return false
		}
	}
	left := true
	if i > start {
		left = VerifySquenceOfBSTRecursive(nums, start, i-1)
	}
	right := true
	if i < end-1 {
		right = VerifySquenceOfBSTRecursive(nums, i+1, end-1)
	}

	return (left && right)
}

func main() {
	fmt.Println(VerifySquenceOfBST([]int{7, 4, 6, 5}))
}
