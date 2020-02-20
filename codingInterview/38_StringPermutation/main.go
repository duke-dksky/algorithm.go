package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}

	rows, cols := len(matrix), len(matrix[0])
	var startX, startY int
	res := make([]int, 0)

	for startX <= rows-1/2 && startY <= cols-1/2 {
		tmp := spiralOrderCore(matrix, startX, startY)
		res = append(res, tmp...)
		startX++
		startY++
	}

	return res
}

func spiralOrderCore(matrix [][]int, startX, startY int) []int {
	if len(matrix) == 0 {
		return nil
	}
	rows, cols := len(matrix), len(matrix[0])
	endX, endY := rows-1-startX, cols-1-startY

	res := make([]int, 0)

	for i := startY; i <= endY; i++ {
		res = append(res, matrix[startX][i])
	}

	if startX < endX {
		for i := startX + 1; i <= endX; i++ {
			res = append(res, matrix[i][endY])
		}
	}

	if startX < endX && startY < endY {
		for i := endY - 1; i >= startY; i-- {
			res = append(res, matrix[endX][i])
		}
	}

	if startY < endY && startX < endY {
		for i := endX - 1; i >= startX+1; i-- {
			res = append(res, matrix[i][startY])
		}
	}
	return res
}

func main() {
	fmt.Println(spiralOrder([][]int{{5, 6, 7}}))
}
