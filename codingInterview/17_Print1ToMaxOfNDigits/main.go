package main

import "fmt"

func PrintNumber(s string) {
	pos := 0
	for pos = 0; pos < len(s); pos++ {
		if s[pos] != '0' {
			break
		}
	}
	if pos == len(s) {
		fmt.Println("0")
	} else {
		fmt.Println(s[pos:])
	}
}

func Print1ToMaxOfNDigits(n int) {
	if n < 0 {
		return
	}
	bs := make([]byte, n)
	Print1ToMaxOfNDigitsRecursively(bs, 0)
}

func Print1ToMaxOfNDigitsRecursively(bs []byte, index int) {
	if index == len(bs) {
		PrintNumber(string(bs))
		return
	}

	for i := 0; i < 10; i++ {
		bs[index] = byte(i) + '0'
		Print1ToMaxOfNDigitsRecursively(bs, index+1)
	}

}

func main() {
	Print1ToMaxOfNDigits(2)

	//fmt.Println(byte(3) + '0')
}
