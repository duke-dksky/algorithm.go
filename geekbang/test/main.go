package main

import "fmt"

func cal(n int) int {
	i, sum := 0, 0
	for i <= n {
		sum += i
		i++
	}
	return sum
}

func main() {
	fmt.Println(cal(100))
}
