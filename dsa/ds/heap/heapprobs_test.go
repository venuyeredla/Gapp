package heap

import "testing"

func TestSorting(t *testing.T) {
	//t.Skip()
	//input := GenArray(5, 20)
	input := []int{10, 80, 30, 90, 40, 50, 70}
	//expected := []int{10, 30, 40, 50, 70, 80, 90}
	HeapSort(input)
}
