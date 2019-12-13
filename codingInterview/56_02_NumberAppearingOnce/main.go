package main

import (
	"fmt"
)

const (
	INT_SIZE = 64
)

func FindNumberAppearingOnce(data []int) int {

	if len(data) == 0 {
		return 0
	}

	bitSum := make([]int, INT_SIZE, INT_SIZE)
	for _, v := range data {
		bitMask := 1
		for i := INT_SIZE - 1; i >= 0; i-- {
			if (v & bitMask) != 0 {
				bitSum[i] += 1
			}
			bitMask = bitMask << 1
		}
	}

	var ret int = 0
	for _, v := range bitSum {
		ret = ret << 1
		ret += v % 3
	}
	return ret
}

func main() {
	fmt.Println(FindNumberAppearingOnce([]int{-10, 214, 214, 214}))
}
