package main

import "fmt"

func hammingWeight(num uint32) int {
	oneCount := 0
	for num > 0 {
		if num&1 == 1 {
			oneCount++
		}
		num >>= 1
	}
	return oneCount
}

func main() {
	fmt.Println(hammingWeight(12))
}
