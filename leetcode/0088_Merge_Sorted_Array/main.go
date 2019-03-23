package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {

	idx := m + n - 1
	for m > 0 && n > 0 {
		if nums1[m-1] > nums2[n-1] {
			nums1[idx] = nums1[m-1]
			m--
		} else {
			nums1[idx] = nums2[n-1]
			n--
		}
		idx--
	}
	for n > 0 {
		nums1[idx] = nums2[n-1]
		n--
		idx--
	}

}

func main() {

	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}

	merge(nums1, 3, nums2, 3)
	fmt.Println(nums1)
}
