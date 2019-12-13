package main

import (
	"fmt"
	"strconv"
)

func Numberof1(n int) int {
	num := 0
	for n > 0 {
		if n%10 == 1 {
			num += 1
		}
		n /= 10
	}
	return num
}

func NumberOf1Between1AndN_1(n int) int {
	num := 0
	for i := 0; i <= n; i++ {
		num += Numberof1(i)
	}
	return num
}

func PowerBase10(n int) int {
	ret := 1
	for i := 0; i < n; i++ {
		ret *= 10
	}
	return ret
}

func NumberOf1Between1AndN(n int) int {
	strN := fmt.Sprint(n)
	length := len(strN)

	firstDigit := int(strN[0] - '0')
	if length == 1 && firstDigit == 0 {
		return 0
	}
	if length == 1 && firstDigit > 0 {
		return 1
	}

	numFirstDigit := 0
	if firstDigit > 1 {
		numFirstDigit = PowerBase10(length - 1)
	} else if firstDigit == 1 {
		tmpVal, _ := strconv.Atoi(strN[1:])
		numFirstDigit = tmpVal + 1
	}
	numOtherDigit := firstDigit * (length - 1) * PowerBase10(length-2)

	tmpVal, _ := strconv.Atoi(strN[1:])
	numRecursive := NumberOf1Between1AndN(tmpVal)

	return numFirstDigit + numOtherDigit + numRecursive
}

func main() {
	fmt.Println(NumberOf1Between1AndN(28))
}
