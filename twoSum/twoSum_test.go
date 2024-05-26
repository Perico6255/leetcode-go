package twosum

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
  var simulations = []struct{
    input []int
    target int
    espected []int
  }{
      {[]int{2,7,11,15},9,[]int{0,1}},
      {[]int{3,2,4},6,[]int{1,2}},
      {[]int{3,3},6,[]int{0,1}},
    }
  for _,s :=range simulations{
    output := TwoSum(s.input, s.target )
    if !reflect.DeepEqual(s.espected, output) {
      t.Errorf("Error, \n EXPECTED %v \n OUTPUT %v",s.espected,output)
    }

  }
  
}
