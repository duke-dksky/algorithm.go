package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x == 0 {
		return true
	}
	var (
		tmp int = x
		y   int = 0
	)
	for tmp != 0 {
		y = y*10 + tmp%10
		tmp = tmp / 10
	}

	if y == x {
		return true
	}
	return false

}

func main() {
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(-121))
	fmt.Println(isPalindrome(10))
}
