package main

import "fmt"

func getRow(rowIndex int) []int {

	ret := make([]int, rowIndex+1)
	for idx, _ := range ret {
		ret[idx] = 1
	}

	for i := 0; i < rowIndex+1; i++ {
		for j := i - 1; j >= 1; j-- {
			ret[j] = ret[j] + ret[j-1]
		}
	}

	return ret

}

func main() {
	fmt.Println(getRow(0))
	fmt.Println(getRow(1))
	fmt.Println(getRow(2))
	fmt.Println(getRow(3))
	fmt.Println(getRow(4))
}
