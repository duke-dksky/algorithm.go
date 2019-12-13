package main

import "fmt"

func FindFirstBitIs1(num int) uint {
	var indexBit uint
	for num&1 == 0 && indexBit < 32 {
		num = num >> 1
		indexBit++
	}
	return indexBit
}

func IsBit1(num int, indexBit uint) bool {
	return ((num >> indexBit) & 1) == 1
}

func FindNumsAppearOnce(data []int) []int {
	ret := make([]int, 2)
	if len(data) == 0 {
		return ret
	}

	resEor := 0
	for _, v := range data {
		resEor ^= v
	}

	indexFirstBitIs1 := FindFirstBitIs1(resEor)

	for _, v := range data {
		if IsBit1(v, indexFirstBitIs1) {
			ret[0] ^= v
		} else {
			ret[1] ^= v
		}
	}
	return ret
}

func main() {
	fmt.Println("vim-go")
}
