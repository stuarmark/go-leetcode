package main

import (
	"testing"
)

// TestAddTwoNumbers - from go-leetcode adding two linked lists of integers
func TestAddTwoNumbers(t *testing.T) {

	numarr1 := [3]int{2, 4, 3}
	numarr2 := [3]int{5, 6, 4}

	l1c := ListNode{Val: numarr1[2], Next: nil}
	l1b := ListNode{Val: numarr1[1], Next: &l1c}
	l1 := ListNode{Val: numarr1[0], Next: &l1b}

	l2c := ListNode{Val: numarr2[2], Next: nil}
	l2b := ListNode{Val: numarr2[1], Next: &l2c}
	l2 := ListNode{Val: numarr2[0], Next: &l2b}

	retNode := addTwoNumbers(&l1, &l2)
	arr, carryover := 2, 0

	for {
		sum := numarr1[arr] + numarr2[arr] + carryover

		if sum > 9 && arr != 0 {
			sum = sum - 10
			carryover = 1
		} else {
			carryover = 0
		}

		if retNode.Val != sum {
			t.Error("retNode.Val: ", retNode.Val, " and sum not equal: ", sum)
		}

		if retNode.Next != nil {
			retNode = retNode.Next
		} else {
			break
		}
		arr--
	}
}

// testcase linked lists w/ non-equivalent lengths
// numarr1 := [2]int{9, 8}
// numarr2 := [1]int{1}
// l1b := ListNode{Val: numarr1[1], Next: nil}
// l1 := ListNode{Val: numarr1[0], Next: &l1b}
// l2 := ListNode{Val: numarr2[0], Next: nil}

// testcase single digit w/ carry over
// l1 := ListNode{Val: 5, Next: nil}
// l2 := ListNode{Val: 5, Next: nil}
