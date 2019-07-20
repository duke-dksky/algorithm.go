package main

import (
	"fmt"
)

// base
func binarySearch(n []int, target int) int {
	begin, end := 0, len(n)-1
	for begin <= end {
		mid := begin + (end-begin)>>1
		if n[mid] == target {
			return mid
		} else if n[mid] < target {
			begin = mid + 1
		} else {
			end = mid - 1
		}
	}
	return -1
}

// 第一个符合条件的值
func binarySearch_1(n []int, target int) int {
	begin, end := 0, len(n)-1
	for begin <= end {
		mid := begin + (end-begin)>>1
		if n[mid] > target {
			end = mid - 1
		} else if n[mid] < target {
			begin = mid + 1
		} else {
			if mid == 0 || n[mid-1] != target {
				return mid
			} else {
				end = mid - 1
			}
		}
	}
	return -1
}

// 第一个大于目标值
func binarySearch_2(n []int, target int) int {
	begin, end := 0, len(n)-1
	for begin <= end {
		mid := begin + (end-begin)>>1
		if n[mid] >= target {
			if mid == 0 || n[mid-1] < target {
				return mid
			} else {
				end = mid - 1
			}
		} else if n[mid] < target {
			begin = mid + 1
		}
	}
	return -1
}

// 循环有序数组
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
		// 无论mid落在何处,左右两边肯定有一边是有序的,有序的边可以判断target是否存在,进而决定边界
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
	n := []int{1, 3, 4, 5, 6, 8, 8, 8, 11, 18}
	fmt.Println(binarySearch_1(n, 8))
}
