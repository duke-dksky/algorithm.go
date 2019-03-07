package main

import "fmt"

func twoSum(nums []int, target int) []int {
	result := make([]int, 2)
	if nums == nil || len(nums) < 2 {
		return result
	}
	tmp := make(map[int]int, len(nums))
	for index, num := range nums {
		if i, exist := tmp[target-num]; exist {
			result[0], result[1] = i, index
		} else {
			tmp[num] = index
		}
	}

	return result
}

func main() {
	nums := []int{2, 7, 11, 15}
	fmt.Println(twoSum(nums, 9))
}
