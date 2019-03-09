package main

import "fmt"

func longestConsecutive(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	mp := make(map[int]bool, len(nums))
	for _, num := range nums {
		mp[num] = true
	}

	maxL := 1
	for k, v := range mp {
		if !v {
			continue
		}
		curL := 1
		toRight, toLeft := true, true
		for i := 1; mp[k+i] || mp[k-i]; i++ {
			if mp[k+i] && toRight {
				curL++
				mp[k+i] = false
			} else {
				toRight = false
			}
			if mp[k-i] && toLeft {
				curL++
				mp[k-i] = false
			} else {
				toLeft = false
			}
		}
		maxL = maxNum(maxL, curL)
	}

	return maxL
}

func maxNum(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{100, 4, 200, 1, 3, 2}
	nums1 := []int{9, 1, 4, 7, 3, -1, 0, 5, 8, -1, 6}
	fmt.Println(longestConsecutive(nums))
	fmt.Println(longestConsecutive(nums1))
}
