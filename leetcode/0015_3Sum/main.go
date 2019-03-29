package main

import (
	"fmt"
	"sort"
)

func threeSum_force(nums []int) [][]int {
	ret := [][]int{}
	if nums == nil || len(nums) < 3 {
		return ret
	}
	sort.Ints(nums)
	n := len(nums)

	for i := 0; i < n; i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < n; j++ {
			if j != i+1 && nums[j] == nums[j-1] {
				continue
			}
			for k := j + 1; k < n; k++ {
				if k != j+1 && nums[k] == nums[k-1] {
					continue
				}
				if nums[i]+nums[j]+nums[k] == 0 {
					ret = append(ret, []int{nums[i], nums[j], nums[k]})
				}
			}
		}
	}
	return ret

}

func threeSum(nums []int) [][]int {
	ret := [][]int{}
	if nums == nil || len(nums) < 3 {
		return ret
	}
	sort.Ints(nums)
	n := len(nums)

	for i := 0; i < n; i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		j, k := i+1, n-1
		for j < k {
			if nums[i]+nums[j]+nums[k] == 0 {
				ret = append(ret, []int{nums[i], nums[j], nums[k]})
				j++
				for j < k && nums[j] == nums[j-1] {
					j++
				}
			} else if nums[i]+nums[j]+nums[k] < 0 {
				j++
			} else {
				k--
			}
		}
	}
	return ret
}

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(nums))
}
