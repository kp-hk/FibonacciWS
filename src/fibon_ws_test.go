// unit test for Calculate_sequence

package main

import (
  "testing"
  "reflect"
)

func TestCalculate_sequence(t *testing.T) {
  expected := []int {0,1,1,2,3,5,8}
  actual := Calculate_sequence(7) 
  if !reflect.DeepEqual(actual, expected) {
    t.Error("Expected", expected)
    t.Error("Actual", actual)
  }
}
