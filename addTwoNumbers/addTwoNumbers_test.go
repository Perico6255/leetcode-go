package addtwonumbers

import (
	"fmt"
	"testing"
)

// func TestAddTwoNumbers(t *testing.T) {
// 	n3 := &ListNode{Val: 3}
// 	n2 := &ListNode{Val: 4, Next: n3}
// 	n1 := &ListNode{Val: 2, Next: n2}
// 	nl3 := &ListNode{Val: 4}
// 	nl2 := &ListNode{Val: 6, Next: nl3}
// 	nl1 := &ListNode{Val: 5, Next: nl2}
//
// 	output := AddTwoNumbers(n1, nl1)
//
// 	if 7!=output.Val {
// 		t.Errorf("Error, \n EXPECTED %v \n OUTPUT %v", 7, output.Val)
// 	}
//   outputNumber:= NodeToInt(output)
//   if 708 !=*outputNumber{
// 		t.Errorf("Error, \n EXPECTED %v \n OUTPUT %v", 708, *outputNumber)
//   }
//
// }

func TestErrorCase(t *testing.T) {
	n2 := &ListNode{Val: 6}
	n1 := &ListNode{Val: 5, Next: n2}
  num := NodeToInt(n1)
  fmt.Println(num)
	nl3 := &ListNode{Val: 9}
	nl2 := &ListNode{Val: 4, Next: nl3}
	nl1 := &ListNode{Val: 5, Next: nl2}
  num2 := NodeToInt(nl1)
  fmt.Println(*num2)
  if  65 != *num {
		t.Errorf("Error, \n EXPECTED %v \n OUTPUT %v", 65, *num)
  } 
  if  945 != *num2 {
		t.Errorf("Error, \n EXPECTED %v \n OUTPUT %v", 945, *num2)
  } 
  numFinal := *num + *num2
  fmt.Println(numFinal)
}

// TODO: the error is if the number end with 0, when we turn the number around we have problems with "0"

func TestNodeToInt(t *testing.T) {
	n3 := &ListNode{Val: 3}
	n2 := &ListNode{Val: 4, Next: n3}
	n1 := &ListNode{Val: 2, Next: n2}
	output := NodeToInt(n1) 	
  if *output != 342 {
		t.Errorf("Error, \n EXPECTED %v \n OUTPUT %v", 342, *output)
	}

	nl3 := &ListNode{Val: 5}
	nl2 := &ListNode{Val: 6, Next: nl3}
	nl1 := &ListNode{Val: 4, Next: nl2}

	output2 := NodeToInt(nl1) 	
  if *output2 != 564 {
		t.Errorf("Error, \n EXPECTED %v \n OUTPUT %v", 564, *output2)
	}
}
func TestIntToNode(t *testing.T) {
  l := IntToNode(*ReverseInt(123))
  fmt.Println(l.Val)

	output := NodeToInt(l) 	

  if *output != 123 {
		t.Errorf("Error, \n EXPECTED %v \n OUTPUT %v", 123, *output)
	}

  
}
