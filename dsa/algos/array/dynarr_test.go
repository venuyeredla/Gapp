package array

import "testing"

func TestSubsetSum(t *testing.T) {
	input := []int{3, 2, 7, 1}
	result := SubsetSum(input, 6)
	resp := SubsetSumD(input, 6)
	if result != resp {
		t.Error("Failed to pass test")
		t.FailNow()
	}
}
