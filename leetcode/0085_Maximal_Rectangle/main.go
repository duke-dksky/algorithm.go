package main

import "fmt"

func maximalRectangle(matrix [][]byte) int {
	if matrix == nil || len(matrix) == 0 {
		return 0
	}
	heights := make([]int, len(matrix[0]))
	maxArea := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == '0' {
				heights[j] = 0
			} else {
				heights[j] = heights[j] + 1
			}
		}
		maxArea = maxNum(maxArea, largestRectangleArea(heights))
	}
	return maxArea

}

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
	matrix1 := [][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}
	matrix2 := [][]byte{
		{'0', '1'},
		{'1', '0'},
	}
	fmt.Println(maximalRectangle(matrix1))
	fmt.Println(maximalRectangle(matrix2))
}
