package main

import "fmt"

func PrintMatrixClockwisely(nums [][]int) {
	rows := len(nums)
	if rows == 0 {
		return
	}
	cols := len(nums[0])
	if cols == 0 {
		return
	}

	start := 0
	for rows > start*2 && cols > start*2 {
		PrintMatricInCircle(nums, start, rows, cols)
		start++
	}
}

func PrintMatricInCircle(nums [][]int, start, rows, cols int) {
	endX := cols - 1 - start
	endY := rows - 1 - start

	for i := start; i <= endX; i++ {
		fmt.Printf("%d ", nums[start][i])
	}

	for i := start + 1; i <= endY; i++ {
		fmt.Printf("%d ", nums[i][endX])
	}

	if endY > start {
		for i := endX - 1; i >= start; i-- {
			fmt.Printf("%d ", nums[endY][i])
		}
	}

	if endX > start {
		for i := endY - 1; i >= start+1; i-- {
			fmt.Printf("%d ", nums[i][start])
		}
	}

}

func main() {
	nums := [][]int{
		{1, 2, 3, 4, 1},
		{4, 5, 6, 7, 1},
		{8, 9, 10, 11, 1},
	}
	PrintMatrixClockwisely(nums)
}
