package array

import (
	"fmt"
	"testing"
)

func AssertEquals(expected, actual []int, log bool) (result bool, message string) {
	if log {
		Printable(expected, 0, len(expected)-1)
		Printable(actual, 0, len(expected)-1)
	}
	if len(expected) == len(actual) {
		for i := 0; i < len(expected); i++ {
			if expected[i] != actual[i] {
				return false, "Failed at index - " + fmt.Sprint(i) + " Expected = " + fmt.Sprint(expected[i]) + " Actual = " + fmt.Sprint(actual[i])
			}
		}
		return true, ""
	} else {
		return false, "Array lengths are unequal"
	}
}

func TestSorting(t *testing.T) {
	//t.Skip()
	//input := GenArray(5, 20)
	input := []int{10, 80, 30, 90, 40, 50, 70}
	expected := []int{10, 30, 40, 50, 70, 80, 90}
	sortAlgos := []Salgo{Bubble, Selection, Insertion, Heap, Quick, Merge}
	for _, algo := range sortAlgos {
		copy := input[:]
		Sort(copy, algo)
		fmt.Printf("Using sort algorithm : %v \n", algo)
		result, msg := AssertEquals(expected, copy, false)
		if !result {
			t.Errorf(msg)
			break
		}
	}
}

func TestArraySearch(t *testing.T) {
	//t.Skip()
	fmt.Println("Linear search")
	input := []int{10, 80, 30, 90, 40, 50, 70}
	key := 30
	found, index := linearSearch(input, key)
	if !found {
		t.Errorf("Erron in linear search algorithm ")
	} else {
		t.Logf("Found at index = %v", index)
	}
	Sort(input, Bubble)
	fmt.Println("Binary search")
	found = binarySearch(input, key)
	if !found {
		t.Fatal("Input: ", input, " Key : ", key, found)
	} else {
		fmt.Println("Input: ", input, " Key : ", key, found)
	}
}

func TestSubset(t *testing.T) {
	arr := []int{1, 2, 3, 4}
	SubArraysR(arr, 0, 0, 4)
	Permuations(arr, 0, len(arr))
	Subset(arr)
	//combinations(arr, 1, 2)
}

func TestLargest(t *testing.T) {
	arr := []int{2, -3, 4, -1, -2, 1, 5, -3}
	LargestSumContiguous(arr)
}

func TestMoveAllzeros(t *testing.T) {
	arr := []int{1, 0, 2, 0, 0, 3}
	expected := []int{1, 2, 3, 0, 0, 0}
	MovallZeros(arr)
	result, msg := AssertEquals(expected, arr, false)
	if !result {
		t.Errorf(msg)
		t.Fail()
	}

}

func TestRearrang(t *testing.T) {
	arr := []int{-1, -1, 6, 1, 9, 3, 2, -1, 4, -1}
	expected := []int{-1, 1, 2, 3, 4, -1, 6, -1, -1, 9}
	Rearrange(arr)
	result, msg := AssertEquals(expected, arr, false)
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
