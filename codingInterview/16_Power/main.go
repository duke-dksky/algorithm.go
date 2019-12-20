package main

import "fmt"

func Power(base, exponent int) float64 {
	if base == 0 && exponent < 0 {
		return 0
	}
	absExponent := exponent
	if exponent < 0 {
		absExponent = -exponent
	}
	result := PowerWithUnsignedExponent(base, absExponent)
	if exponent < 0 {
		result = 1.0 / result
	}

	return result
}

func PowerWithUnsignedExponent(base, exponent int) float64 {
	if exponent == 0 {
		return 1
	}
	if exponent == 1 {
		return float64(base)
	}
	result := PowerWithUnsignedExponent(base, exponent>>1)
	result *= result
	if exponent&0x1 == 1 {
		result *= float64(base)
	}
	return result
}

func main() {
	fmt.Println("vim-go")
}
