package main

import "fmt"

func Permutation(s string) {
	if len(s) == 0 {
		return
	}
	PermutationRecursive([]byte(s), 0)
}

func PermutationRecursive(b []byte, begin int) {
	if len(b) == begin {
		fmt.Println(string(b))
	}
	for i := begin; i < len(b); i++ {
		b[i], b[begin] = b[begin], b[i]
		PermutationRecursive(b, begin+1)
		b[begin], b[i] = b[i], b[begin]
	}
}

func main() {
	Permutation("abc")
}
