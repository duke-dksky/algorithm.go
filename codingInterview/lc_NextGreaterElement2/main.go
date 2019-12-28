package main

import "fmt"

func NextGreaterElements(nums []int) []int {
	n := len(nums)
	if n == 0 {
		return nil
	}
	res := make([]int, n)
	for i := range res {
		res[i] = -1
	}
	for i := range nums {
		for j := i + 1; j < i+n; j++ {
			if nums[j%n] > nums[i] {
				res[i] = nums[j%n]
				break
			}
		}
	}
	return res
}

func NextGreaterElements_2(nums []int) []int {
	n := len(nums)
	if n == 0 {
		return nil
	}
	res := make([]int, n)
	for i := range res {
		res[i] = -1
	}
	stack := make([]int, 0)
	for i := 0; i < 2*n; i++ {
		num := nums[i%n]
		for len(stack) != 0 && nums[stack[len(stack)-1]] < num {
			res[stack[len(stack)-1]] = num
			stack = stack[:len(stack)-1]
		}
		if i < n {
			stack = append(stack, i)
		}
	}
	return res
}

func main() {
	fmt.Println(NextGreaterElements_2([]int{1, 1, 1, 2, 1}))
}
