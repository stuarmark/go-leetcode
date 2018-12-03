package main

import (
	"fmt"
)

func main() {
	x := 10001
	fmt.Println(isPalindrome(x))
}

func isPalindrome(x int) bool {
	if x < 0 || x%10 == 0 {
		return false
	} else if x <= 9 {
		return true
	}

	pnd := 0
	for x > pnd {
		x = x / 10
		rem := x % 10
		pnd = pnd*10 + rem

		if x == pnd || x/10 == pnd {
			return true
		}
	}
	return false
}
