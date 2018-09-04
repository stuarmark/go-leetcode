package main

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	numchars := 7
	str := "abcaeiou"

	retnumber := lengthOfLongestSubstring(str)

	if retnumber != numchars {
		t.Error("String: ", str, " has ", numchars, " but has returned ", retnumber, " characters")
	}
}
