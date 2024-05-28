package addtwonumbers

import (
	"reflect"
	"testing"
)
func TestAddTwoNumbers(t *testing.T) {
  var simulations = []struct{
    l1, l2, espected ListNode
  }{
    }
  for _,s :=range simulations{
    // output := TwoSum(s.input, s.target )
    // if !reflect.DeepEqual(s.espected, output) {
    //   t.Errorf("Error, \n EXPECTED %v \n OUTPUT %v",s.espected,output)
    // }
    output := AddTwoNumbers(&s.l1, &s.l2)
    if !reflect.DeepEqual(s.espected, output) {
      t.Errorf("Error, \n EXPECTED %v \n OUTPUT %v",s.espected,output)
    }

  }
  
}

