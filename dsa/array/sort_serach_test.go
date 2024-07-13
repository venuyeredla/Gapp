package array

import (
	"Gapp/dsa/utils"
	"fmt"
	"testing"
)

func TestSortBsearch(t *testing.T) {
	//t.Skip()
	input := []int{10, 80, 30, 90, 40, 50, 70}
	expected := []int{10, 30, 40, 50, 70, 80, 90}
	sortAlgos := []Salgo{Bubble, Selection, Insertion, Merge, Quick}
	fmt.Println("Sorting and applying binary search")
	for _, algo := range sortAlgos {
		C := make([]int, len(input))
		copy(C, input)
		Sort(C, algo)
		result, msg := utils.AssertEquals(expected, C, false)
		if !result {
			t.Errorf("Failed Algorithm : %v Error MSG=%v ", algo, msg)
			break
		}
	}
}

func TestFindkey(t *testing.T) {
	index := findKeyRotated([]int{5, 1, 3}, 5)
	expected := 0
	if index != expected {
		t.Errorf("Expected = %v , Actual =%v ", expected, index)
	}
}

func TestMoveAllzeros(t *testing.T) {
	arr := []int{1, 0, 2, 0, 0, 3}
	expected := []int{1, 2, 3, 0, 0, 0}
	MovallZeros(arr)
	result, msg := utils.AssertEquals(expected, arr, false)
	if !result {
		t.Errorf(msg)
		t.Fail()
	}

}

func TestRearrang(t *testing.T) {
	arr := []int{-1, -1, 6, 1, 9, 3, 2, -1, 4, -1}
	expected := []int{-1, 1, 2, 3, 4, -1, 6, -1, -1, 9}
	Rearrange(arr)
	result, msg := utils.AssertEquals(expected, arr, false)
	if !result {
		t.Errorf(msg)
		t.Fail()
	}
}

func TestKmost(t *testing.T) {
	arr := []int{3, 1, 4, 4, 5, 2, 6, 1, 2}
	KMostOccurance(arr, 2)
}

func TestRemoveDuplicates(t *testing.T) {
	arr := []int{2, 3, 5, 5, 7, 11, 11, 11, 13}
	removeDuplicates(arr)
}

func TestRotation(t *testing.T) {
	/*arr := []int{1, 2, 3, 4, 5, 6, 7}
	Rotation(arr, 3)
	fmt.Println(arr) */
	arr := []int{2, 0}
	Jump(arr)
}
