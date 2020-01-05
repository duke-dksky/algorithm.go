package main

import "fmt"

func Find(matrix [][]int, rows int, columns int, number int) bool {
	ret := false

	if matrix != nil && rows > 0 && columns > 0 {
		row := 0
		column := columns - 1
		for column >= 0 && row < rows {
			if matrix[row][column] > number {
				column -= 1
			} else if matrix[row][column] < number {
				row += 1
			} else {
				ret = true
				break
			}
		}
	}

	return ret
}

func main() {
	matrix := [][]int{
		{1, 2, 8, 9},
		{2, 4, 9, 12},
		{4, 7, 10, 13},
		{6, 8, 11, 15},
	}
	if Find(matrix, 4, 4, 16) {
		fmt.Println("Find")
	} else {
		fmt.Println("not Find")
	}
}
