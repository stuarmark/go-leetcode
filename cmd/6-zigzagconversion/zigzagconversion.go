package main

import (
	"fmt"
)

func main() {
	str := `P   A   H   N
A P L S I I G
Y   I   R`

	fmt.Println(convert(str, 3))
}

func convert(s string, numRows int) string {
	if numRows <= 1 {
		return s
	}

	return s
}
