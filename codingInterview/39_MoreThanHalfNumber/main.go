package main

//import "fmt"

func Partition(data []int, start, end int) int {
	return 0
}

func MoreThanHalfNum(data []int) int {
	if len(data) == 0 {
		return 0
	}
	num := data[1]
	count := 1
	for i := 1; i < len(data); i++ {
		if num == data[i] {
			count++
		} else {
			count--
			if count == 0 {
				num = data[i]
				count = 1
			}
		}
	}
	if CheckMoreThanHalf(data, num) {
		return num
	}
	return 0
}

func CheckMoreThanHalf(data []int, num int) bool {
	countNum := 0
	for _, v := range data {
		if v == num {
			countNum++
		}
	}
	if len(data) < 2*countNum {
		return true
	}
	return false
}
