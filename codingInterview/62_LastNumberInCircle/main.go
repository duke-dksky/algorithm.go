package main

import "fmt"

func LastRemaining(n, m int) int {
	if m < 1 || n < 1 {
		return -1
	}
	last := 0
	for i := 1; i <= n; i++ {
		last = (last + m) % i
	}
	return last
}

func main() {
	fmt.Println("vim-go")
}
