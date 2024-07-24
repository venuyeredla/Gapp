package linked

import (
	"container/list"
	"fmt"
)

func goList() {
	list := list.New()
	list.PushBack(10)
	list.PushBack(12)

	list.PushFront(24)

	temp := list.Front()

	for temp != nil {
		fmt.Printf("%v  ", temp.Value)
		temp = temp.Next()
	}

	//list.Back()
	//list.Front()
	//list.PushBack()

}

/**  Problems   **/

/*
# Notes
Use sentinel(dummy head)

# Problems.
1. Merge two sorted lists.
2. Reverse SLL & DLL.
3. Cycles in SLL.
4. Test for overlappeing lists.
5. Remove duplicates from sorted list.
6. Cyclic right shift
7. Test for palindrome of list.
8. Add list based integers.
*/

/*
Trying to merge into first tree.
*/
func AddTwoNumbers(list1, list2 *list.List) *list.List {
	c1 := list1.Front()
	c2 := list2.Front()
	carry := 0
	for c1 != nil && c2 != nil {
		num1 := c1.Value.(int)
		num2 := c2.Value.(int)
		sum := num1 + num2 + carry
		if sum > 9 {
			diggit := sum % 10
			carry = sum / 10
			c1.Value = diggit
		} else {
			c1.Value = sum
		}
		c1 = c1.Next()
		c2 = c2.Next()
	}
	return list1
}
