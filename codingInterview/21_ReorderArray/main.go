package main

import (
	"fmt"
)

func ReorderOddEven(data []int) {
	if len(data) == 0 {
		return
	}
	Reorder(data, isEven)
}

func Reorder(data []int, f func(int) bool) {
	if len(data) == 0 {
		return
	}
	start, end := 0, len(data)-1
	for start < end {
		for start < end && f(data[start]) {
			start++
		}
		for start < end && !f(data[end]) {
			end--
		}
		data[start], data[end] = data[end], data[start]
	}
}

func isEven(num int) bool {
	return (num & 0x1) == 0
}

func main() {
	data := []int{1, 2, 3, 4, 5}
	ReorderOddEven(data)
	fmt.Println(data)
}
