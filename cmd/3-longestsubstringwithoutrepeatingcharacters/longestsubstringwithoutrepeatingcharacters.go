package main

import (
	"fmt"
)

func main() {
	str := "abcaeiou"
	//str := "dvdf"
	//str := "ckcilbkd"
	fmt.Println(lengthOfLongestSubstring(str))
}

func lengthOfLongestSubstring(s string) int {
	ln, max := 0, 0
	reftable := [256]int{}
	for i := range reftable {
		reftable[i] = -1
	}
	for i := 0; i < len(s); i++ {
		fmt.Println(i)
		fmt.Println(reftable[s[i]])
		fmt.Println(ln)
		if reftable[s[i]] >= ln {
			ln = reftable[s[i]] + 1
		}
		reftable[s[i]] = i
		if max < i-ln+1 {
			max = i - ln + 1
		}
	}
	return max
}
