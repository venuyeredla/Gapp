package array

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func AssertEquals(expected, actual []int, log bool) bool {
	if log {
		Printable(expected, 0, len(expected)-1)
		Printable(actual, 0, len(expected)-1)
	}

	if len(expected) == len(actual) {
		for i := 0; i < len(expected); i++ {
			if expected[i] != actual[i] {
				return false
			}
		}
		return true
	} else {
		return false
	}
}

func TestSorting(t *testing.T) {
	//t.Skip()
	//input := GenArray(5, 20)
	input := []int{10, 80, 30, 90, 40, 50, 70}
	expected := []int{10, 30, 40, 50, 70, 80, 90}
	salgo := Merge
	Sort(input, salgo)
	if !AssertEquals(expected, input, false) {
		t.Errorf("Sorting not right")
	}
}

func TestLinearSearch(t *testing.T) {
	t.Skip()
	fmt.Println("Linear search")
	randArray := GenArray(10, 50)
	rand.Seed(time.Now().UnixMilli())
	key := rand.Intn(20)

	//found := linearSearch(randArray, key)
	Sort(randArray, Insertion)
	fmt.Println("Binary search")
	found := binarySearch(randArray, key)
	if !found {
		t.Fatal("Input: ", randArray, " Key : ", key, found)
	} else {
		fmt.Println("Input: ", randArray, " Key : ", key, found)
	}
}

func TestSubArrays(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	RSubArrays(arr, 0, 0, 5)
}

func TestSubset(t *testing.T) {
	arr := []int{1, 2, 3}
	//combinations(arr, 1, 2)
	//Subset(arr)
	Permuations(arr, 0, len(arr))
}

func TestLargest(t *testing.T) {
	arr := []int{2, -3, 4, -1, -2, 1, 5, -3}
	LargestSumContiguous(arr)
}

func TestMoveAllzeros(t *testing.T) {
	arr := []int{1, 0, 2, 0, 0, 3}
	expected := []int{1, 2, 3, 0, 0, 0}
	MovallZeros(arr)
	if !AssertEquals(expected, arr, false) {
		t.Errorf("Sorting not right")
		t.Fail()
	}
}

func TestRearrang(t *testing.T) {
	arr := []int{-1, -1, 6, 1, 9, 3, 2, -1, 4, -1}
	expected := []int{-1, 1, 2, 3, 4, -1, 6, -1, -1, 9}

	Rearrange(arr)
	if !AssertEquals(expected, arr, true) {
		t.Errorf("Sorting not right")
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
