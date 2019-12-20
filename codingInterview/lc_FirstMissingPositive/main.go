package main

import "fmt"

func FirstMissingPositive(data []int) int {
	if len(data) == 0 {
		return 0
	}

	for i := range data {
		for data[i] > 0 && data[i] <= len(data) && data[i] != data[data[i]-1] {
			data[i], data[data[i]-1] = data[data[i]-1], data[i]
		}
	}
	for i := range data {
		if data[i] != i+1 {
			return i + 1
		}
	}
	return len(data) + 1
}

func main() {
	fmt.Println(FirstMissingPositive([]int{3, 4, -1, 1}))
}
