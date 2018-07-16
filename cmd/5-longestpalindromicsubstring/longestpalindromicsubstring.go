package main

import "fmt"

func main() {
	s := "babad"

	fmt.Println(longestPalindrome(s))
}

func longestPalindrome(s string) string {
	l, ml, j, k := 0, 0, 0, 0
	if len(s) <= 1 {
		return s
	}
	for i := 0; i < len(s); i++ {
		j, k = i, i
		for k >= 0 && j < len(s) && s[k] == s[j] {
			k--
			j++
		}
		if j-k+1-2 > ml {
			ml = j - k + 1 - 2
			l = k + 1
		}

		j, k = i+1, i

		for k >= 0 && j < len(s) && s[k] == s[j] {
			k--
			j++
		}
		if j-k+1-2 > ml {
			ml = j - k + 1 - 2
			l = k + 1
		}
	}
	return string(s[l : l+ml])
}
