package main

import "fmt"

func NumberOf1_Solution1(n int) int {
	flag := 1
	count := 0
	for flag != 0 {
		if n&flag != 0 {
			count++
		}
		flag = flag << 1
	}
	return count
}

func NumberOf1_Solution2(n int) int {
	count := 0
	for n > 0 {
		count++
		n = (n - 1) & n
	}
	return count
}
func main() {
	fmt.Println(NumberOf1_Solution1(10))
}
