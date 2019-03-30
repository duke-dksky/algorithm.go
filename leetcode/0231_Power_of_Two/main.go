package main

import "fmt"

func isPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	}
	oneCount := 0
	for n > 0 {
		if n&1 == 1 {
			oneCount++
		}
		n >>= 1
	}
	if oneCount <= 1 {
		return true
	}
	return false
}

func main() {
	fmt.Println(isPowerOfTwo(161))
}
