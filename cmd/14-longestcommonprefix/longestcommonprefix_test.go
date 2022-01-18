package main

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	strArr := []string{"florida", "flowing", "jumpin"}

	strMatch := longestCommonPrefix(strArr)

	if strMatch != "flo" {
		t.Error("String returned: ", strMatch, " does not match flo")
	}
}
