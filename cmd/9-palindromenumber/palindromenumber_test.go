package main

import (
	"fmt"
	"testing"
)

func TestPalindromeNumber(t *testing.T) {

	firstInt := 1001
	firstBoolCheck := true
	firstBool := isPalindrome(firstInt)

	if firstBool == firstBoolCheck {
		fmt.Printf("The firstInt %d is a palindrome \n", firstInt)
	}

	secondInt := 1023
	secondBoolCheck := true
	secondBool := isPalindrome(secondInt)

	if secondBool != secondBoolCheck {
		fmt.Printf("The secondInt %d is not a palindrome \n", secondInt)
	}
}
