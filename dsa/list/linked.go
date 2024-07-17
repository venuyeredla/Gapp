package list

import (
	"Gapp/dsa/errors"
	"Gapp/dsa/types"
	"fmt"
	"strings"
)

// A doubly linked list node.
type Node struct {
	Data       types.Hashable
	Next, Prev *Node
}

// Compares the Data of the node to the passed element.
func (n *Node) Equals(b types.Equatable) bool {
	switch x := b.(type) {
	case *Node:
		return n.Data.Equals(x.Data)
	default:
		return n.Data.Equals(b)
	}
}

// Compares the Data of the node to the passed element.
func (n *Node) Less(b types.Sortable) bool {
	switch x := b.(type) {
	case *Node:
		return n.Data.Less(x.Data)
	default:
		return n.Data.Less(b)
	}
}

// Hashes the Data of the node to the passed element.
func (n *Node) Hash() int {
	return n.Data.Hash()
}

// A doubly linked list. There is no synchronization.
// The fields are publically accessible to allow for easy customization.
type LinkedList struct {
	Length int
	Head   *Node
	Tail   *Node
}

func New() *LinkedList {
	return &LinkedList{
		Length: 0,
		Head:   nil,
		Tail:   nil,
	}
}

func (l *LinkedList) Size() int {
	return l.Length
}

func (l *LinkedList) Items() (it types.KIterator) {
	cur := l.Head
	it = func() (item types.Hashable, _ types.KIterator) {
		if cur == nil {
			return nil, nil
		}
		item = cur.Data
		cur = cur.Next
		return item, it
	}
	return it
}

func (l *LinkedList) Backwards() (it types.KIterator) {
	cur := l.Tail
	it = func() (item types.Hashable, _ types.KIterator) {
		if cur == nil {
			return nil, nil
		}
		item = cur.Data
		cur = cur.Prev
		return item, it
	}
	return it
}

func (l *LinkedList) Has(item types.Hashable) bool {
	for x, next := l.Items()(); next != nil; x, next = next() {
		if x.Equals(item) {
			return true
		}
	}
	return false
}

func (l *LinkedList) Push(item types.Hashable) (err error) {
	return l.EnqueBack(item)
}

func (l *LinkedList) Pop() (item types.Hashable, err error) {
	return l.DequeBack()
}

func (l *LinkedList) EnqueFront(item types.Hashable) (err error) {
	n := &Node{Data: item, Next: l.Head}
	if l.Head != nil {
		l.Head.Prev = n
	} else {
		l.Tail = n
	}
	l.Head = n
	l.Length++
	return nil
}

func (l *LinkedList) EnqueBack(item types.Hashable) (err error) {
	n := &Node{Data: item, Prev: l.Tail}
	if l.Tail != nil {
		l.Tail.Next = n
	} else {
		l.Head = n
	}
	l.Tail = n
	l.Length++
	return nil
}

func (l *LinkedList) DequeFront() (item types.Hashable, err error) {
	if l.Head == nil {
		return nil, errors.Errorf("List is empty")
	}
	item = l.Head.Data
	l.Head = l.Head.Next
	if l.Head != nil {
		l.Head.Prev = nil
	} else {
		l.Tail = nil
	}
	l.Length--
	return item, nil
}

func (l *LinkedList) DequeBack() (item types.Hashable, err error) {
	if l.Tail == nil {
		return nil, errors.Errorf("List is empty")
	}
	item = l.Tail.Data
	l.Tail = l.Tail.Prev
	if l.Tail != nil {
		l.Tail.Next = nil
	} else {
		l.Head = nil
	}
	l.Length--
	return item, nil
}

func (l *LinkedList) First() (item types.Hashable) {
	if l.Head == nil {
		return nil
	}
	return l.Head.Data
}

func (l *LinkedList) Last() (item types.Hashable) {
	if l.Tail == nil {
		return nil
	}
	return l.Tail.Data
}

/*
// Can be compared to any types.IterableContainer
func (l *LinkedList) Equals(b types.Equatable) bool {
	if o, ok := b.(types.IterableContainer); ok {
		return list.Equals(l, o)
	} else {
		return false
	}
}

// Can be compared to any types.IterableContainer
func (l *LinkedList) Less(b types.Sortable) bool {
	if o, ok := b.(types.IterableContainer); ok {
		return list.Less(l, o)
	} else {
		return false
	}
}

func (l *LinkedList) Hash() int {
	return list.Hash(l)
}
*/

func (l *LinkedList) String() string {
	if l.Length <= 0 {
		return "{}"
	}
	items := make([]string, 0, l.Length)
	for item, next := l.Items()(); next != nil; item, next = next() {
		items = append(items, fmt.Sprintf("%v", item))
	}
	return "{" + strings.Join(items, ", ") + "}"
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
