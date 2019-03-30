package main

import "fmt"

func missingNumber(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return -1
	}

	ret, n := 0, len(nums)
	for i := 0; i <= n; i++ {
		ret = ret ^ i
	}
	for _, num := range nums {
		ret = ret ^ num
	}
	return ret

}

func main() {
	nums := []int{0, 1, 3}
	fmt.Println(missingNumber(nums))
}
