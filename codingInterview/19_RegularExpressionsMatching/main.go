package main

import "fmt"

func Match(str, pattern string) bool {
	// pattern is empty
	if len(str) == 0 && len(pattern) == 0 {
		return true
	}
	if len(str) != 0 && len(pattern) == 0 {
		return false
	}
	return MatchCore([]byte(str), []byte(pattern))
}

func MatchCore(bStr, bPattern []byte) bool {
	if len(bStr) == 0 && len(bPattern) == 0 {
		return true
	}
	if len(bStr) != 0 && len(bPattern) == 0 {
		return false
	}

	if len(bPattern) > 1 && bPattern[1] == '*' {
		if len(bStr) != 0 && (bStr[0] == bPattern[0] || bPattern[0] == '.') {
			return MatchCore(bStr[1:], bPattern[2:]) || MatchCore(bStr[1:], bPattern) || MatchCore(bStr, bPattern[2:])
		} else {
			return MatchCore(bStr, bPattern[2:])
		}
	}

	if len(bStr) != 0 && (bStr[0] == bPattern[0] || bPattern[0] == '.') {
		return MatchCore(bStr[1:], bPattern[1:])
	}

	return false
}

func main() {
	fmt.Println("vim-go")
}
