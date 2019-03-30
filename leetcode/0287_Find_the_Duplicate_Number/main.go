package main

import "fmt"

func findDuplicate(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return -1
	}
	ret, n := 0, len(nums)
	for i := 0; i <= n-1; i++ {
		ret = ret ^ i
	}
	for _, num := range nums {
		ret = ret ^ num
	}
	return ret
}

func main() {
	nums := []int{3, 1, 3, 4, 2}
	fmt.Println(findDuplicate(nums))
}
