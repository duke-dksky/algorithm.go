package main

import "fmt"

func BuildProductionArray(data []int) []int {
	if len(data) < 2 {
		return nil
	}

	ret := make([]int, len(data), len(data))

	// 前半部分
	ret[0] = 1
	for i := 1; i < len(data); i++ {
		ret[i] = data[i-1] * ret[i-1]
	}

	// 后半部分
	temp := 1
	for i := len(data) - 2; i >= 0; i-- {
		temp *= data[i+1]
		ret[i] *= temp
	}
	return ret

}

func main() {
	fmt.Println(BuildProductionArray([]int{1, 2, 3, 4, 5}))
	fmt.Println(BuildProductionArray([]int{1, 2, 0, 4, 5}))
	fmt.Println(BuildProductionArray([]int{1, -2, 3, -4, 5}))
}
