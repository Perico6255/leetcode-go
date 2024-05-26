package twosum

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
  input  := []int{2,7,11,15}
  var target int = 9
  expected  := []int{0,1}
  output := TwoSum(input, target )
  if !reflect.DeepEqual(expected, output) {
    t.Errorf("Error, \n EXPECTED %v \n OUTPUT %v",expected,output)
    
  }
  
}
