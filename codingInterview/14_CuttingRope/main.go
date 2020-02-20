package main

import (
	"fmt"
	"math"
)

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func maxProductAfterCutting_2(n int) int {
	if n < 2 {
		return 0
	}
	if n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}
	res := make([]int, n+1)
	res[0], res[1], res[2], res[3] = 0, 1, 2, 3

	for i := 4; i <= n; i++ {
		max := 0
		for j := 1; j <= n/2; j++ {
			product := res[i-j] * res[j]
			if max < product {
				max = product
			}
		}
		res[i] = max
	}
	return res[n+1]
}

func maxProductAfterCutting_1(n int) int {
	if n < 2 {
		return 0
	}
	if n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}

	timesOf3 := n / 3
	if n-timesOf3*3 == 1 {
		timesOf3 -= 1
	}

	timesOf2 := (n - timesOf3*3) / 2

	return int(pow(3, timesOf3) * pow(2, timesOf2))
}

func main() {
	fmt.Println(maxProductAfterCutting_1(50))
}
