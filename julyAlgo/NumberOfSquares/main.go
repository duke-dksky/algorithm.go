package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func printArray(a [][]int) {
	for k, _ := range a {
		fmt.Printf("%v\n", a[k])
	}
}

func NumberOfSquares(n, m int) int {
	ret := 0
	array := make([][]int, m)
	for k, _ := range array {
		array[k] = make([]int, n)
	}
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			array[i][j] = min(i, j)
			ret += array[i][j]
		}
	}
	printArray(array)
	return ret
}

func main() {
	fmt.Println(NumberOfSquares(2, 2))
	fmt.Println(NumberOfSquares(19, 19))
}
