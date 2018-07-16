package main

import (
	"fmt"
)

func main() {
	a := "abab"
	b := "abab"
	stringswap := buddyStrings(a, b)

	fmt.Println(fmt.Sprintf("%s AND %s produce: %t", a, b, stringswap))
}

func buddyStrings(A string, B string) bool {
	lenA, lenB := len(A), len(B)
	if lenA == 0 || lenB == 0 || lenA != lenB {
		return false
	}

	if A == B {
		s := make(map[rune]int, lenA)
		for i, char := range A {
			s[char] = i
		}
		return len(s) < lenA
	}

	aset, j, i := make([]int, 2), 0, 0
	for ; i < lenA; i++ {
		if A[i] != B[i] && j < 2 {
			aset[j] = i
			j++
		}
	}

	if j > 0 {
		return (A[aset[1]] == B[aset[0]] && A[aset[0]] == B[aset[1]])
	}
	return (A[0] == A[1])
}
