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

func TestLargest(t *testing.T) {
	arr := []int{2, -3, 4, -1, -2, 1, 5, -3}
	LargestSumSubArray(arr)
}

func TestLargeMonotonic(t *testing.T) {
	arr := []int{3, 10, 2} // 3, 10, 2, 1, 20
	result := longIncreasingSubseq(arr, 0, 0)
	if result != 3 {
		t.Error("Failed to pass test")
		t.FailNow()
	}
}
