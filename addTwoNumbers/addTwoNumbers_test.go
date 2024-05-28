package addtwonumbers

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	n3 := &ListNode{Val: 3}
	n2 := &ListNode{Val: 4, Next: n3}
	n1 := &ListNode{Val: 2, Next: n2}
	nl3 := &ListNode{Val: 5}
	nl2 := &ListNode{Val: 6, Next: nl3}
	nl1 := &ListNode{Val: 4, Next: nl2}
	fmt.Println(nl3.Next)

	output := AddTwoNumbers(n1, nl1)

	if !reflect.DeepEqual(7, &output.Val) {
		t.Errorf("Error, \n EXPECTED %v \n OUTPUT %v", 7, output.Val)
	}
}
func TestNodeToList(t *testing.T) {
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
