package main

import "fmt"

func longestSubstringWithoutDuplication(s string) int {
	position := make([]int, 26)
	for i := range position {
		position[i] = -1
	}
	curLength := 0
	maxLength := 0
	for i, c := range s {
		preIndex := position[int(c-'a')]
		if preIndex < 0 || i-preIndex > curLength {
			curLength++
		} else {
			curLength = i - preIndex
		}
		if curLength > maxLength {
			maxLength = curLength
		}
		position[int(c-'a')] = i
	}
	return maxLength
}

func main() {
	fmt.Println("vim-go")
}
