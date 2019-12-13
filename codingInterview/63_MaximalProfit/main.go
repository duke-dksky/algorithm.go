package main

import "fmt"

func MaxDiff(data []int) int {
	if len(data) < 2 {
		return 0
	}

	min := data[0]
	if data[1] < data[0] {
		min = data[1]
	}
	maxDiff := data[1] - data[0]

	for i := 2; i < len(data); i++ {
		curDiff := data[i] - min
		if curDiff > maxDiff {
			maxDiff = curDiff
		}
		if data[i] < min {
			min = data[i]
		}
	}

	return maxDiff
}

func main() {
	fmt.Println(MaxDiff([]int{4, 1, 3, 2, 5}))
}
