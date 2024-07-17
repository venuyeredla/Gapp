package array

import (
	"Gapp/dsa/utils"
	"testing"
)

func TestPrdoductExceptItself(t *testing.T) {
	input := []int{-1, 1, 0, -3, 3} // 1, 0, 2
	expected := []int{0, 0, 9, 0, 0}
	output := productExceptSelf(input)
	utils.AssertEquals(expected, output, false)
}

func TestDuplicate(t *testing.T) {
	input := []int{1, 3, 4, 2, 2}
	expected := 2
	output := findDuplicate(input)
	if output != expected {
		t.Errorf("Expected = %v and Actual =%v", expected, output)
		t.FailNow()
	}
}
