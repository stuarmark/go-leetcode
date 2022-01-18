package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(longestCommonPrefix([]string{"florida", "flowing", "jumpin"}))
}

// 1 <= strs.length <= 200
// 0 <= strs[i].length <= 200
// strs[i] consists of only lower-case English letters.
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	} else if len(strs) == 1 {
		return strs[0]
	}

	strB := strings.Builder{}

	first := strs[0]
	rest := strs[1:]

	for i := 0; i < len(first); i++ {
		j := first[i]
		for _, w := range rest {
			if len(w) < i+1 {
				return strB.String()
			}
			if w[i] != j {
				return strB.String()
			}
		}
		strB.WriteByte(j)
	}
	return strB.String()
}
