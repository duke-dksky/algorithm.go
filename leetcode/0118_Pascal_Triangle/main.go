package main

import "fmt"

func generate(numRows int) [][]int {

	switch numRows {
	case 0:
		return [][]int{}
	case 1:
		return [][]int{{1}}
	case 2:
		return [][]int{{1}, {1, 1}}
	}

	ret := [][]int{{1}, {1, 1}}
	for k := 3; k <= numRows; k++ {
		tmp := make([]int, k)
		ret = append(ret, tmp)
		ret[k-1][0], ret[k-1][k-1] = 1, 1

		for n := 1; n < k-1; n++ {
			ret[k-1][n] = ret[k-2][n-1] + ret[k-2][n]
		}

	}

	return ret
}

func main() {
	fmt.Println(generate(0))
	fmt.Println(generate(1))
	fmt.Println(generate(2))
	fmt.Println(generate(3))
	fmt.Println(generate(4))
	fmt.Println(generate(5))
}
