package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	ret := [][]int{}
	if nums == nil || len(nums) < 3 {
		return ret
	}
	sort.Ints(nums)

}

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(nums))
}
