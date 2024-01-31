package array

import (
	"Gapp/dsa/util"
	"fmt"
	"testing"
)

func TestSortBsearch(t *testing.T) {
	//t.Skip()
	input := []int{10, 80, 30, 90, 40, 50, 70}
	expected := []int{10, 30, 40, 50, 70, 80, 90}
	key := 70
	sortAlgos := []Salgo{Bubble, Selection, Insertion, Merge, Quick}
	fmt.Println("Sorting and applying binary search")
	for _, algo := range sortAlgos {
		C := make([]int, len(input))
		copy(C, input)
		Sort(C, algo)
		result, msg := util.AssertEquals(expected, C, false)
		if !result {
			t.Errorf("Failed Algorithm : %v Error MSG=%v ", algo, msg)
			break
		}
		found, index := BinarySearch(C, key)
		if !found {
			t.Fatal("Input: ", input, " Key : ", key, found)
		} else {
			fmt.Printf("Algorithm = %v , Sorted =%v - Fount at=%v \n", algo, C, index)
		}
	}
}

func TestSubset(t *testing.T) {
	arr := []int{1, 2, 3, 4}
	subsets := Subset(arr)
	for _, v := range subsets {
		fmt.Println(v)
	}
	//combinations(arr, 1, 2)
}

func TestSubset2(t *testing.T) {
	arr := []int{1, 2, 3} //, 3, 4
	SubsetBackTracking(arr, 0, 2)
}

func TestMoveAllzeros(t *testing.T) {
	arr := []int{1, 0, 2, 0, 0, 3}
	expected := []int{1, 2, 3, 0, 0, 0}
	MovallZeros(arr)
	result, msg := util.AssertEquals(expected, arr, false)
	if !result {
		t.Errorf(msg)
		t.Fail()
	}

}

func TestRearrang(t *testing.T) {
	arr := []int{-1, -1, 6, 1, 9, 3, 2, -1, 4, -1}
	expected := []int{-1, 1, 2, 3, 4, -1, 6, -1, -1, 9}
	Rearrange(arr)
	result, msg := util.AssertEquals(expected, arr, false)
	if !result {
		t.Errorf(msg)
		t.Fail()
	}
}

func TestKmost(t *testing.T) {
	arr := []int{3, 1, 4, 4, 5, 2, 6, 1}
	KMostOccur(arr)
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
