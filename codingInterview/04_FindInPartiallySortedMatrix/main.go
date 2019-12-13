package main

import "fmt"

func searchRange(nums []int, target int) []int {
	if nums == nil || len(nums) == 0 {
		return []int{}
	}
	ret := []int{-1, -1}
	i, j := 0, len(nums)-1
	for i <= j {
		mid := i + (j-i)>>2
		if nums[mid] >= target {
			j = mid - 1
		} else {
			i = mid + 1
		}
	}
	if i < len(nums) && nums[i] == target {
		ret[0] = i
	} else {
		return ret
	}

	j = len(nums) - 1
	for i <= j {
		mid := i + (j-i)>>2
		if nums[mid] <= target {
			i = mid + 1
		} else {
			j = mid - 1
		}
	}
	ret[1] = j
	return ret

}

func main() {
	nums := []int{5, 7, 7, 8, 8, 10}
	fmt.Println(searchRange(nums, 8))
}
