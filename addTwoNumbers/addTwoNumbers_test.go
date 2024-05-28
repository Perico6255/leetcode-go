package addtwonumbers

import (
	"reflect"
	"testing"
)
func TestAddTwoNumbers(t *testing.T) {
  n3 := &ListNode{Val:2}
  n2 := &ListNode{Val:4, Next: n3}
  n1 := &ListNode{Val:3, Next: n2}
  nl3 := &ListNode{Val:4}
  nl2 := &ListNode{Val:6, Next: nl3}
  nl1 := &ListNode{Val:5, Next: nl2}

  output:= AddTwoNumbers(n1, nl1)
  

  if !reflect.DeepEqual(7,&output.Val ) {
    t.Errorf("Error, \n EXPECTED %v \n OUTPUT %v", 7,&output.Val)
  }


  
}

