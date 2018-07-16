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

	// testcase linked lists w/ equivalent lengths
	numarr1 := [3]int{2, 4, 3}
	numarr2 := [3]int{5, 6, 4}

	l1c := ListNode{Val: numarr1[2], Next: nil}
	l1b := ListNode{Val: numarr1[1], Next: &l1c}
	l1 := ListNode{Val: numarr1[0], Next: &l1b}

	l2c := ListNode{Val: numarr2[2], Next: nil}
	l2b := ListNode{Val: numarr2[1], Next: &l2c}
	l2 := ListNode{Val: numarr2[0], Next: &l2b}

	// testcase linked lists w/ non-equivalent lengths
	// numarr1 := [2]int{9, 8}
	// numarr2 := [1]int{1}
	// l1b := ListNode{Val: numarr1[1], Next: nil}
	// l1 := ListNode{Val: numarr1[0], Next: &l1b}
	// l2 := ListNode{Val: numarr2[0], Next: nil}

	// testcase singel digit w/ carry over
	// l1 := ListNode{Val: 5, Next: nil}
	// l2 := ListNode{Val: 5, Next: nil}

	retNode := addTwoNumbers(&l1, &l2)
	testLinkedList(retNode)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	a, b, carryover := 0, 0, 0
	listNodeRet := &ListNode{Val: ((l1.Val + l2.Val) % 10)}
	headNode := listNodeRet

	if (l1.Val + l2.Val) > 9 {
		carryover = 1
	}

	for {
		if l1.Next != nil {
			l1 = l1.Next
		} else {
			a = 1
			l1.Val = 0
		}
		if l2.Next != nil {
			l2 = l2.Next
		} else {
			b = 1
			l2.Val = 0
		}
		if a == 1 && b == 1 {
			break
		}

		ltemp := l1.Val + l2.Val + carryover
		node := ListNode{Val: ltemp % 10, Next: nil}
		listNodeRet.Next = &node

		if ltemp > 9 {
			carryover = 1
		} else {
			carryover = 0
		}
		listNodeRet = listNodeRet.Next
	}
	if carryover == 1 {
		node := ListNode{Val: carryover, Next: nil}
		listNodeRet.Next = &node
	}

	return headNode
}

func testLinkedList(list *ListNode) {
	for {
		fmt.Println(list)
		if list.Next != nil {
			list = list.Next
		} else {
			break
		}
	}
}
