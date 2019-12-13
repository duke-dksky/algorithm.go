package main

import "fmt"

func Sum(n int) int {
	if n == 1 {
		return 1
	}
	return n + Sum(n-1)
}

func main() {
	fmt.Println("vim-go")
}
