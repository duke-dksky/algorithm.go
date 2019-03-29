package main

import (
	"fmt"
	"sort"
)

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func threeSumClosest(nums []int, target int) int {
	if nums == nil || len(nums) < 3 {
		return -1
	}
	sort.Ints(nums)
	n := len(nums)

	ret := nums[0] + nums[1] + nums[2]
	for i := 0; i < n-2; i++ {
		j, k := i+1, n-1
		for j < k {
			current := nums[i] + nums[j] + nums[k]
			if current == target {
				return target
			} else if current < target {
				j++
			} else {
				k--
			}
			if abs(current-target) < abs(ret-target) {
				ret = current
			}
		}
	}
	return ret

}

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSumClosest(nums, 1))
}
