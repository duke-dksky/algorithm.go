package main

import "fmt"

func maxSlidingWindow(nums []int, k int) []int {
	queue := make([]int, 0)
	for i := 0; i < k; i++ {
		for len(queue) != 0 && nums[queue[len(queue)-1]] < nums[i] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, i)
	}

	ret := make([]int, 0)
	ret = append(ret, nums[queue[0]])

	for i := k; i < len(nums); i++ {
		for len(queue) != 0 && nums[queue[len(queue)-1]] < nums[i] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, i)
		if i-queue[0] >= k {
			queue = queue[1:]
		}
		ret = append(ret, nums[queue[0]])
	}
	return ret
}

func main() {
	fmt.Println(maxSlidingWindow([]int{1, -1}, 1))
}
