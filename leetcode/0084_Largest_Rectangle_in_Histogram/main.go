package main

import "fmt"

func largestRectangleArea(heights []int) int {
	stack := make([]int, 0)
	i, max := 0, 0

	for i < len(heights) {
		if len(stack) == 0 || heights[i] > heights[stack[len(stack)-1]] {
			stack = append(stack, i)
			i++
		} else {
			pop := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			h := heights[pop]
			var w int
			if len(stack) == 0 {
				w = i
			} else {
				w = i - 1 - stack[len(stack)-1]
			}
			max = maxNum(max, h*w)
		}
	}

	for len(stack) != 0 {
		pop := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		h := heights[pop]
		var w int
		if len(stack) == 0 {
			w = i
		} else {
			w = i - 1 - stack[len(stack)-1]
		}
		max = maxNum(max, h*w)
	}

	return max
}

func maxNum(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	heights := []int{2, 1, 5, 6, 2, 3}
	fmt.Println(largestRectangleArea(heights))
}
