package main

import "fmt"

func maxNum(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func minNum(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxArea(height []int) int {
	result := 0

	if height == nil || len(height) < 2 {
		return 0
	}
	begin, end := 0, len(height)-1
	for begin < end {
		result = maxNum(result, (end-begin)*minNum(height[end], height[begin]))
		if height[end] < height[begin] {
			end--
		} else {
			begin++
		}
	}

	return result
}

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(height))
}
