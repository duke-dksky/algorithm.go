package main

import "fmt"

func ReplaceBlank(str []byte) []byte {

	if str == nil {
		return str
	}

	originLen := len(str)
	for i := 0; i < originLen; i++ {
		if str[i] == ' ' {
			str = append(str, ' ')
			str = append(str, ' ')
		}
	}
	newLen := len(str)

	indexOfOriginal := originLen - 1
	indexOfNew := newLen - 1
	for indexOfOriginal >= 0 && indexOfNew > indexOfOriginal {
		fmt.Printf("%c\n", str[indexOfOriginal])
		if str[indexOfOriginal] == ' ' {
			str[indexOfNew] = '0'
			indexOfNew--
			str[indexOfNew] = '2'
			indexOfNew--
			str[indexOfNew] = '%'
			indexOfNew--
		} else {
			str[indexOfNew] = str[indexOfOriginal]
			indexOfNew--
		}
		indexOfOriginal--
	}

	return str
}

func main() {
	str := []byte(" yang dong kui ")
	str = ReplaceBlank(str)
	fmt.Printf("%c\n", str)
}
