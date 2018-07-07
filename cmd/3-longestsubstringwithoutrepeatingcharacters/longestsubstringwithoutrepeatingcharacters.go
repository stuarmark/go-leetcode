package main

import (
	"fmt"
)

func main() {
	str := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring(str))
}

func lengthOfLongestSubstring(s string) int {
	mr := make(map[rune]int)
	ln, max := 0, 0
	for i := 0; i < len(s); i++ {
		mr[rune(s[i])] = i
		if i > 0 && mr[rune(s[i])] > max {
			max = mr[rune(s[i])]
			fmt.Println(fmt.Sprintf("max:%d", max))
		}

		tot := (i - max + 1)
		fmt.Println(fmt.Sprintf("i-max+1: %d", tot))
		if (i - max + 1) > ln {
			ln = (i - max + 1)
			fmt.Println(fmt.Sprintf("ln: %d", ln))
		}
	}
	fmt.Println(mr)
	return ln
}
