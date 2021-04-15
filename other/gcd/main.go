package main

import "fmt"

func gcd(m, n int) int {
	if m%n == 0 {
		return n
	}
	return gcd(n, m%n)
}

func main() {
	fmt.Println(gcd(3, 7))
}
