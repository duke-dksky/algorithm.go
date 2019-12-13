package main

import "fmt"

func Add(n1, n2 int) int {
	carry := (n1 & n2) << 1
	sum := n1 ^ n2

	for carry != 0 {
		n1, n2 = sum, carry
		carry = (n1 & n2) << 1
		sum = n1 ^ n2
	}
	return sum

}

func main() {
	fmt.Println("vim-go")
}
