package main

import "fmt"

func plusOne(digits []int) []int {
	one := 1
	for i := len(digits) - 1; i >= 0; i-- {
		num := digits[i] + one
		one = num / 10
		digits[i] = num % 10
		if one == 0 {
			break
		}
	}
	if one == 1 {
		return append([]int{1}, digits...)
	}
	return digits
}

func main() {
	input1 := []int{1, 2, 3}
	input2 := []int{9, 9}
	fmt.Println(plusOne(input1))
	fmt.Println(plusOne(input2))
}
