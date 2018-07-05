package main

import (
	"fmt"
)

// ListNode -  Main struct for addtwonumbers
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	numarr1 := [3]int{2, 4, 3}
	numarr2 := [3]int{5, 6, 4}

	l1c := ListNode{Val: numarr1[2], Next: nil}
	l1b := ListNode{Val: numarr1[1], Next: &l1c}
	l1 := ListNode{Val: numarr1[0], Next: &l1b}

	l2c := ListNode{Val: numarr2[2], Next: nil}
	l2b := ListNode{Val: numarr2[1], Next: &l2c}
	l2 := ListNode{Val: numarr2[0], Next: &l2b}

	retNode := addTwoNumbers(&l1, &l2)

	fmt.Println(retNode)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	const dec = 0
	a, b := 0, 0

	for {
		fmt.Println(fmt.Sprintf("l1 is: %d", l1.Val))
		fmt.Println(fmt.Sprintf("l2 is: %d", l2.Val))
		if l1.Next != nil {
			l1 = l1.Next
		} else {
			a = 1
		}
		if l2.Next != nil {
			l2 = l2.Next
		} else {
			b = 1
		}

		if a == 1 && b == 1 {
			break
		}
	}
	return nil
}
