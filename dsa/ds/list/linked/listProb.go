package linked

import "Gapp/dsa/ds/types"

/*
# Notes
Use sentinel(dummy head)

# Problems.
1. Merge two sorted lists.
2. Reverse SLL & DLL.
3. Cycles in SLL.
4. Test for overlappeing lists.
5. Remove duplicates form sorted list.
6. Cyclic right shift
7. Test for palindrome of list.
8. Add list based integers.
*/
func MergeSorted(list1, list2 *LinkedList) *LinkedList {

	return nil
}

/*
Trying to merge into first tree.
*/
func AddTwoNumbers(list1, list2 *LinkedList) *LinkedList {

	head1 := list1.Head
	head2 := list2.Head

	carry := 0
	for head1 != nil && head2 != nil {
		num1 := head1.Data.(types.Int)
		num2 := head2.Data.(types.Int)
		sum := int(num1+num2) + carry
		if sum > 9 {
			diggit := sum % 10
			carry = sum / 10
			head1.Data = types.Int(diggit)
		} else {
			head1.Data = types.Int(sum)
		}
		head1 = head1.Next
		head2 = head2.Next
	}
	return list1
}

/*
At each what do you do?

	P - C  - next
*/
func ReverseSLL(list1 *LinkedList) *LinkedList {
	current := list1.Head
	var prev *Node
	for current != nil {
		temp := current.Next
		current.Next = prev
		prev = current
		current = temp
	}
	list1.Head = prev
	return list1
}

func RemoveDuplicatesSorted(list1 *LinkedList) *LinkedList {

	return nil
}

func RemoveDuplicatesUnSorted(list1 *LinkedList) *LinkedList {

	return nil
}
