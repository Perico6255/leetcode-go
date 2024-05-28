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

	return l1
}
func NodeToInt(l *ListNode)*int{
  var n int
  for i := 1; true; i *= 10 {
    n = n + l.Val * i
    if l.Next == nil {
      return &n
    }
    l = l.Next
  }
  return &n
}
func IntToNode(n int) *ListNode {
  l := &ListNode{}
  fmt.Println(l)
  for ; n >0 ; n/=10 {
    l.Val = n%10
  }
  
  return l
  
}
