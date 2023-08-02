package binary

import (
	"Gapp/dsa/ds/types"
	"fmt"
	"testing"
)

func TestInsert(t *testing.T) {
	expected := []int{5, 3, 7, 4, 6}
	bst := NewBST()

	for _, i := range expected {
		if !bst.BstPut(types.Int(i), types.Int(i)) {
			t.Errorf("Element %v should have been added to the tree", i)
		}
	}

	for _, i := range expected {
		if bst.Find(types.Int(i)) == nil {
			t.Errorf("Element %v expected to be in the tree, but was not", i)
		}
	}

	if bst.BstPut(types.Int(4), 44) {
		t.Error("Duplicate elements should not be added")
	}

	if bst.Find(types.Int(4)) == 44 {
		t.Error("Previously inserted elements should not be updated")
	}

	if c := bst.count; c != len(expected) {
		t.Errorf("Tree expected to have %v elements, but has %v instead", len(expected), c)
	}
}

func TestRemove_SingleElement(t *testing.T) {
	bst := NewBST()

	bst.BstPut(types.Int(5), 10)

	if !bst.Delete(types.Int(5)) {
		t.Errorf("Element %v should have been removed", 5)
	}

	if bst.Find(types.Int(5)) != nil {
		t.Errorf("Element %v should not have been found", 5)
	}
}

func TestRemove_RootWithSingleChild(t *testing.T) {
	bst := NewBST()

	bst.BstPut(types.Int(5), 10)
	bst.BstPut(types.Int(4), 8)

	if !bst.Delete(types.Int(5)) {
		t.Errorf("Element %v should have been removed", 5)
	}

	if bst.Find(types.Int(5)) != nil {
		t.Errorf("Element %v should not have been found", 5)
	}

	if bst.Find(types.Int(4)) == nil {
		t.Errorf("Element with key %v was not found", 4)
	}

	if bst.count != 1 {
		t.Errorf("Expected element count %v found cound %v", 1, bst.count)
	}
}

func TestRemove_RootWithTwoChildren(t *testing.T) {
	bst := NewBST()

	bst.BstPut(types.Int(5), 10)
	bst.BstPut(types.Int(4), 8)
	bst.BstPut(types.Int(6), 12)

	if !bst.Delete(types.Int(5)) {
		t.Errorf("Element %v should have been removed", 5)
	}

	if bst.Find(types.Int(5)) != nil {
		t.Errorf("Element %v should not have been found", 5)
	}

	if bst.Find(types.Int(4)) == nil {
		t.Errorf("Element with key %v was not found", 4)
	}

	if bst.Find(types.Int(6)) == nil {
		t.Errorf("Element with key %v was not found", 6)
	}

	if bst.count != 2 {
		t.Errorf("Expected element count %v found cound %v", 1, bst.count)
	}
}

func TestRemove(t *testing.T) {
	expected := []int{5, 3, 7, 4, 6}
	bst := NewBST()

	for _, i := range expected {
		bst.BstPut(types.Int(i), i)
	}

	if !bst.Delete(types.Int(6)) {
		t.Errorf("Element %v should have been removed from the tree", 6)
	}

	for _, i := range expected[0:3] {
		if bst.Find(types.Int(i)) == nil {
			t.Errorf("Element %v expected to be in the tree, but was not", i)
		}
	}

	if d := expected[len(expected)-1]; bst.Find(types.Int(d)) != nil {
		t.Errorf("Element %v should have been removed from the tree", d)
	}

	if bst.Delete(types.Int(6)) {
		t.Error("Duplicate elements should not be delete")
	}

	if c := bst.count; c != len(expected)-1 {
		t.Errorf("Tree expected to have %v elements, but has %v instead", len(expected)-1, c)
	}
}

func TestTraverse_InOrder(t *testing.T) {
	elements := []int{5, 3, 7, 4, 6}
	expected := []int{3, 4, 5, 6, 7}

	bst := NewBST()

	for _, i := range elements {
		bst.BstPut(types.Int(i), i)
	}

	i := 0
	for e := range bst.Traverse(InOrder) {
		if e != expected[i] {
			t.Errorf("Expected to traverse %v, but instead traversed %v", expected[i], e)
		}
		i++
	}
}

func TestTraverse_PreOrder(t *testing.T) {
	elements := []int{5, 3, 7, 4, 6}
	expected := []int{5, 3, 4, 7, 6}

	bst := NewBST()

	for _, i := range elements {
		bst.BstPut(types.Int(i), i)
	}

	i := 0
	for e := range bst.Traverse(PreOrder) {
		if e != expected[i] {
			t.Errorf("Expected to traverse %v, but instead traversed %v", expected[i], e)
		}
		i++
	}
}

func TestTraverse_PostOrder(t *testing.T) {
	elements := []int{5, 3, 7, 4, 6}
	expected := []int{4, 3, 6, 7, 5}

	bst := NewBST()

	for _, i := range elements {
		bst.BstPut(types.Int(i), i)
	}

	i := 0
	for e := range bst.Traverse(PostOrder) {
		if e != expected[i] {
			t.Errorf("Expected to traverse %v, but instead traversed %v", expected[i], e)
		}
		i++
	}
}

func TestClear(t *testing.T) {
	elements := []int{5, 3, 7, 4, 6}

	bst := NewBST()

	for _, i := range elements {
		bst.BstPut(types.Int(i), i)
	}

	bst.Clear()

	if c := bst.count; c != 0 {
		t.Errorf("Expected tree to be empty, but has %v elements", c)
	}

	if p := bst.String(); len(p) != 0 {
		t.Errorf("No elements expected in the tree, but found %v", p)
	}
}

func (t *BinaryTree) String() (s string) {
	print(t.root, &s)
	return
}

func print(n *BinaryNode, s *string) {
	if n == nil {
		return
	}

	*s += fmt.Sprintf("%p %v\n", n, n)
	print(n.left, s)
	print(n.right, s)
}

func BenchmarkInsert(b *testing.B) {
	bst := NewBST()
	for _, i := range rand.Perm(b.N) {
		bst.BstPut(types.Int(i), i)
	}
}

func BenchmarkDelete(b *testing.B) {
	bst := NewBST()
	for _, i := range rand.Perm(b.N) {
		bst.BstPut(types.Int(i), i)
	}

	b.ResetTimer()
	for _, i := range rand.Perm(b.N) {
		bst.Delete(types.Int(i))
	}
}

func BenchmarkFind(b *testing.B) {
	bst := NewBST()
	for _, i := range rand.Perm(b.N) {
		bst.BstPut(types.Int(i), i)
	}

	b.ResetTimer()
	for _, i := range rand.Perm(b.N) {
		bst.Find(types.Int(i))
	}
}
