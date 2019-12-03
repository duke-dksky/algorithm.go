package main

import "fmt"

func GetFirstK(data []int, start, end int, k int) int {
	if start > end {
		return -1
	}
	midIndex := start + (end-start)>>1
	midNum := data[midIndex]
	if midNum == k {
		if midIndex > 0 && data[midIndex-1] == k {
			end = midIndex - 1
		} else {
			return midIndex
		}
	} else if midNum > k {
		end = midIndex - 1
	} else if midNum < k {
		start = midIndex + 1
	}
	return GetFirstK(data, start, end, k)
}

func GetLastK(data []int, start, end int, k int) int {
	if start > end {
		return -1
	}
	midIndex := start + (end-start)>>1
	midNum := data[midIndex]
	if midNum == k {
		if midIndex < end && data[midIndex+1] == k {
			start = midIndex + 1
		} else {
			return midIndex
		}
	} else if midNum > k {
		end = midIndex - 1
	} else if midNum < k {
		start = midIndex + 1
	}
	return GetLastK(data, start, end, k)
}

func GetNumberOfK(data []int, k int) int {
	number := 0
	if len(data) == 0 {
		return number
	}
	start, end := 0, len(data)-1
	first := GetFirstK(data, start, end, k)
	last := GetLastK(data, start, end, k)
	if first != -1 && last != -1 {
		number = last - first + 1
	}
	return number
}

func main() {
	fmt.Println(GetNumberOfK([]int{1, 2, 3, 3, 3, 3, 4, 5}, 3))
}
