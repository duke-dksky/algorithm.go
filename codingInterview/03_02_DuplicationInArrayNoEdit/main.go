package main

import "fmt"

func GetDuplication(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	start, end := 1, len(nums)-1
	for start <= end {
		mid := start + (end-start)>>1
		count := countRange(nums, start, mid)

		if start == end {
			if count > 1 {
				return start
			} else {
				break
			}
		}

		if count > mid-start+1 {
			end = mid
		} else {
			start = mid + 1
		}
	}
	return -1
}

func countRange(nums []int, start, end int) int {
	if len(nums) == 0 {
		return 0
	}
	count := 0
	for i := range nums {
		if start <= nums[i] && nums[i] <= end {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(GetDuplication([]int{2, 3, 5, 4, 3, 2, 6, 7}))
	fmt.Println(GetDuplication([]int{3, 2, 1, 4, 4, 5, 6, 7}))
}
