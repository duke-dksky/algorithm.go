package main

import "fmt"

// f(i,j) = max(f(i-1,j), f(i,j-1)) + gift(i,j)

func MaxValue(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func GetMaxValue(values [][]int, rows, cols int) int {
	if rows == 0 && cols == 0 {
		return values[0][0]
	}
	if rows == 0 {
		return values[rows][cols] + GetMaxValue(values, rows, cols-1)
	}
	if cols == 0 {
		return values[rows][cols] + GetMaxValue(values, rows-1, cols)
	}
	return values[rows][cols] + MaxValue(GetMaxValue(values, rows, cols-1), GetMaxValue(values, rows-1, cols))

}

func GetMaxValue_solution1(values [][]int) int {
	if len(values) == 0 {
		return 0
	}
	rows := len(values)
	cols := len(values[0])
	return GetMaxValue(values, rows-1, cols-1)
}

func GetMaxValue_solution2(values [][]int) int {
	if len(values) == 0 {
		return 0
	}
	rows := len(values)
	cols := len(values[0])

	tmp := make([][]int, rows)
	for i := range tmp {
		tmp[i] = make([]int, cols)
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			var up, left int
			if r > 0 {
				up = tmp[r-1][c]
			}
			if c > 0 {
				left = tmp[r][c-1]
			}
			tmp[r][c] = values[r][c] + MaxValue(up, left)
		}
	}
	return tmp[rows-1][cols-1]
}

func GetMaxValue_solution3(values [][]int) int {
	if len(values) == 0 {
		return 0
	}
	rows := len(values)
	cols := len(values[0])

	tmp := make([]int, cols)

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			left, up := 0, 0
			if r > 0 {
				up = tmp[c]
			}
			if c > 0 {
				left = tmp[c-1]
			}
			tmp[c] = MaxValue(up, left) + values[r][c]
		}
	}
	return tmp[cols-1]
}

func main() {
	fmt.Println("vim-go")
}
