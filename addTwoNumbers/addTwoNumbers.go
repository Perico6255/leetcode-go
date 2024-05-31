package addtwonumbers

import "fmt"

/*
You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.
*/

// Definition for single-linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	n1 := NodeToInt(l1)
	n2 := NodeToInt(l2)
	return IntToNode(*ReverseInt(*n1 + *n2))
}
func NodeToInt(l *ListNode) *int {
	var n int
	for i := 1; true; i *= 10 {
		n = n + l.Val*i
		if l.Next == nil {
			return &n
		}
		l = l.Next
	}
	return &n
}
func IntToNode(n int) *ListNode {
	var l *ListNode
	l = &ListNode{Val: n % 10, Next: l}
	n /= 10
	for ; n > 0; n /= 10 {
		l = &ListNode{Val: n % 10, Next: l}
	}
	return l
}
func ReverseInt(n int) *int {
	reversed := 0

	for n > 0 {
		reversed = reversed*10 + n%10
		n /= 10
	}
	return &reversed
}
func IntToNodeReversed(n int) *ListNode {
  l1 := &ListNode{}
  l2 := &l1
  fmt.Println(l1)
  fmt.Println(l2)
	for ; n > 0; n /= 10 {
    l1 = &ListNode{Val: n, Next: *l2}
    l2 = &l1
	}
	return l1
}
