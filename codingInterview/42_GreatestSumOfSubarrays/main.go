package main

import (
	"fmt"
	"math"
)

func FindGreatestSumOfSubArray(data []int) int {
	if len(data) == 0 {
		return 0
	}
	curSum := 0
	greatestSum := math.MinInt64
	for _, v := range data {
		if curSum <= 0 {
			curSum = v
		} else {
			curSum += v
		}
		if greatestSum < curSum {
			greatestSum = curSum
		}
	}
	return greatestSum

}

func main() {
	fmt.Println("vim-go")
}
