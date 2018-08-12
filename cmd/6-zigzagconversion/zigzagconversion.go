package main

import (
	"fmt"
)

func main() {
	str := `PAYPALISHIRING`
	fmt.Print(convert(str, 4))
}

func convert(s string, numRows int) string {
	if numRows == 1 || numRows <= 0 || numRows > len(s) {
		return s
	}
	arr := make([][]byte, numRows)
	k := 0
	incdec := true

	for i := 0; i < len(s); i++ {
		arr[k] = append(arr[k], s[i])

		if incdec {
			k++
			if k == numRows-1 {
				incdec = false
			}
		} else {
			k--
			if k == 0 {
				incdec = true
			}
		}
	}
	retstring := make([]byte, 0)
	for _, str := range arr {
		retstring = append(retstring, str...)
	}
	return string(retstring)
}
