package main

import "fmt"

func findMissingRanges(nums []int, start int, end int) []string {
	ret := make([]string, 0)
	prev := start - 1
	cur := 0
	for i := 0; i <= len(nums); i++ {

		if i != len(nums) {
			cur = nums[i]
		} else {
			cur = end + 1
		}

		if cur-prev >= 2 {
			ret = append(ret, getRange(prev+1, cur-1))
		}
		prev = cur
	}

	return ret

}

func getRange(from, to int) string {
	if from == to {
		return fmt.Sprintf("%v", from)
	}
	return fmt.Sprintf("%v->%v", from, to)
}

func main() {
	fmt.Println(findMissingRanges([]int{0, 1, 3, 50, 75}, 0, 99))
	fmt.Println(findMissingRanges([]int{3, 50, 75}, 0, 99))
}
