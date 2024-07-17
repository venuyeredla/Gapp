package list

import (
	"Gapp/dsa/types"
	"fmt"
	"testing"
)

func TestAddTwo(t *testing.T) {

	list1 := New()
	list2 := New()
	input := []int{2, 4, 3}
	for _, val := range input {
		list1.EnqueBack(types.Int(val))
	}

	input2 := []int{5, 6, 4}
	for _, val := range input2 {
		list2.EnqueBack(types.Int(val))
	}

	result := AddTwoNumbers(list1, list2)
	temp := result.Head
	for temp != nil {
		fmt.Print(temp.Data.(types.Int))
		temp = temp.Next
	}

}

func TestReverse(t *testing.T) {

	list1 := New()
	input := []int{2, 4, 3}
	for _, val := range input {
		list1.EnqueBack(types.Int(val))
	}
	result := ReverseSLL(list1)
	temp := result.Head
	for temp != nil {
		fmt.Print(temp.Data.(types.Int))
		temp = temp.Next
	}
}
