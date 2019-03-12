package main

import "fmt"

func maxNum(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lengthOfLongestSubstring1(s string) int {
	if len(s) == 0 {
		return 0
	}

	maxL := 0

	mp := [128]int{}
	for i, _ := range mp {
		mp[i] = -1
	}

	lowIndex := 0
	for i, c := range s {
		if lowIndex <= mp[c] {
			lowIndex = mp[c] + 1
		} else {
			maxL = maxNum(maxL, i-lowIndex+1)
		}
		mp[c] = i
	}

	return maxL
}

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	maxL := 0

	mp := make(map[rune]int)
	lowIndex := 0
	for i, c := range s {
		if _, exist := mp[c]; exist && lowIndex <= mp[c] {
			lowIndex = mp[c] + 1
		} else {
			maxL = maxNum(maxL, i-lowIndex+1)
		}
		mp[c] = i
	}

	return maxL
}

func main() {
	s1 := "abcabcbb"
	s2 := "bbbbb"
	s3 := "pwwkew"
	s4 := "au"
	s5 := "dvdf"
	s6 := "tmmzuxt"
	fmt.Println(lengthOfLongestSubstring1(s1))
	fmt.Println(lengthOfLongestSubstring1(s2))
	fmt.Println(lengthOfLongestSubstring1(s3))
	fmt.Println(lengthOfLongestSubstring1(s4))
	fmt.Println(lengthOfLongestSubstring1(s5))
	fmt.Println(lengthOfLongestSubstring1(s6))
}
