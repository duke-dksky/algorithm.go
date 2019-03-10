package main

import "fmt"

func trap(height []int) int {
	if height == nil || len(height) < 3 {
		return 0
	}

	ret := 0
	low, high := 0, len(height)-1
	leftMax, rightMax := height[0], height[len(height)-1]
	for low < high {
		if leftMax < rightMax {
			low++
			if leftMax < height[low] {
				leftMax = height[low]
			} else {
				ret += (leftMax - height[low])
			}
		} else {
			high--
			if rightMax < height[high] {
				rightMax = height[high]
			} else {
				ret += (rightMax - height[high])
			}
		}
	}

	return ret
}

func main() {
	height1 := []int{2, 0, 2}
	height2 := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println(trap(height1))
	fmt.Println(trap(height2))
}
