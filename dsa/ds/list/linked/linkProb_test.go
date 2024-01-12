package linked

import (
	"Gapp/dsa/ds/types"
	"fmt"
	"testing"
)

func TestAddTwo(t *testing.T) {

	list1 := New()
	list1.EnqueBack(types.Int(2))
	list1.EnqueBack(types.Int(4))
	list1.EnqueBack(types.Int(3))

	list2 := New()
	list2.EnqueBack(types.Int(5))
	list2.EnqueBack(types.Int(6))
	list2.EnqueBack(types.Int(4))

	result := AddTwoNumbers(list1, list2)

	temp := result.Head
	for temp != nil {
		fmt.Print(temp.Data.(types.Int))
		temp = temp.Next
	}

}

func TestReverse(t *testing.T) {

	list1 := New()
	list1.EnqueBack(types.Int(2))
	list1.EnqueBack(types.Int(4))
	list1.EnqueBack(types.Int(3))

	result := ReverseSLL(list1)
	temp := result.Head
	for temp != nil {
		fmt.Print(temp.Data.(types.Int))
		temp = temp.Next
	}
}
