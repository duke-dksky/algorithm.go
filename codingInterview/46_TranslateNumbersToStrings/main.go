package main

import (
	"fmt"
)

func GetTranslationCount(n int) int {
	if n < 0 {
		return 0
	}
	strN := fmt.Sprint(n)
	count := make([]int, len(strN))
	for i := len(strN) - 1; i >= 0; i-- {
		if i == len(strN)-1 {
			count[i] = 1
		} else {
			count[i] = count[i+1]

			digit1 := int(strN[i] - '0')
			digit2 := int(strN[i+1] - '0')
			converted := digit1*10 + digit2
			if converted >= 10 && converted <= 25 {
				if i < len(strN)-2 {
					count[i] += count[i+2]
				} else {
					count[i] += 1
				}
			}
		}
	}
	return count[0]
}

func main() {
	fmt.Println(GetTranslationCount(125))
}
