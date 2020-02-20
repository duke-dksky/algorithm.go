package main

import "fmt"

func isNumeric(s string) bool {
	length := len(s)
	if length == 0 {
		return false
	}

	integer, dec, exp := false, false, true
	var index int

	index, integer = scanInteger(s, 0)
	if index < length && s[index] == '.' {
		index, dec = scanUnsignedInteger(s, index+1)
	}
	if index < length && (s[index] == 'e' || s[index] == 'E') {
		index, exp = scanInteger(s, index+1)
	}

	return (integer || dec) && exp && (index == length)
}

func scanInteger(s string, start int) (int, bool) {
	if start < len(s) && (s[start] == '-' || s[start] == '+') {
		start++
	}
	return scanUnsignedInteger(s, start)
}

func scanUnsignedInteger(s string, start int) (int, bool) {
	before := start
	for start < len(s) && s[start] >= '0' && s[start] <= '9' {
		start++
	}
	return start, start > before
}

func main() {
	fmt.Println(isNumeric("100"))
}
