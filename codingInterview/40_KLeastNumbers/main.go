package main

import "fmt"

func findKthLargest(nums []int, k int) int {
	heap := nums[:k]
	headSort(heap)
	for i := k; i < len(nums); i++ {
		if heap[0] < nums[i] {
			heap[0] = nums[i]
			adjustHeap(heap, 0, k)
		}
	}
	return heap[k-1]
}

func headSort(nums []int) {
	if len(nums) == 0 {
		return
	}
	for i := len(nums)/2 - 1; i >= 0; i-- {
		adjustHeap(nums, i, len(nums))
	}
	for i := len(nums) - 1; i >= 1; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		adjustHeap(nums, 0, i)
	}
}

func adjustHeap(nums []int, i int, n int) {
	if len(nums) == 0 {
		return
	}
	j := i*2 + 1

	for j < n {
		if j+1 < n && nums[j+1] < nums[j] {
			j += 1
		}
		if nums[i] < nums[j] {
			break
		}
		nums[i], nums[j] = nums[j], nums[i]
		i = j
		j = i*2 + 1
	}
}

func main() {
	nums := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	headSort(nums)
	fmt.Println(nums)
	fmt.Println(findKthLargest(nums, 4))
}
