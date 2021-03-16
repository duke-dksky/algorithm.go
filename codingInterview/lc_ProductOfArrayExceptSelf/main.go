package main

import "fmt"

func productExceptSelf(nums []int) []int {
	ret := make([]int, len(nums))
	ret[0] = 1
	for i := 1; i < len(nums); i++ {
		ret[i] = ret[i-1] * nums[i-1]
	}

	rightProduct := 1
	for i := len(nums) - 2; i >= 0; i-- {
		rightProduct *= nums[i+1]
		ret[i] = ret[i] * rightProduct
	}
	return ret
}

func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
}
