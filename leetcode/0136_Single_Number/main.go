package main

import "fmt"

func singleNumber(nums []int) int {
	ret := 0
	for _, num := range nums {
		ret = ret ^ num
	}
	return ret
}

func main() {
	nums := []int{4, 1, 2, 1, 2}
	fmt.Println(singleNumber(nums))
}
