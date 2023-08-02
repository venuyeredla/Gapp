// Package bst implements an unbalanced binary search tree.
package binary

import "Gapp/dsa/ds/types"

// Insert adds a given key+value to the tree and returns true if it was added.
// Average: O(log(n)) Worst: O(n)

func NewBST() *BinaryTree {
	return &BinaryTree{treeType: BST}
}

func (t *BinaryTree) BstPut(k types.Hashable, v interface{}) (updated bool) {
	t.root, updated = t.root.BstPut(k, v)
	if updated {
		t.count++
	}

	return updated
}

// insert recusively adds a key+value in the tree.
func (n *BinaryNode) BstPut(k types.Hashable, v interface{}) (r *BinaryNode, added bool) {
	if r = n; n == nil {
		// keep track of how many elements we have in the tree
		// to optimize the channel length during traversal
		r = &BinaryNode{key: k, value: v}
		added = true
	} else if k.Less(n.key) {
		r.left, added = n.left.BstPut(k, v)
	} else if n.key.Less(k) {
		r.right, added = n.right.BstPut(k, v)
	}

	return
}

// Delete removes a given key from the tree and returns true if it was removed.
// Average: O(log(n)) Worst: O(n)
func (t *BinaryTree) Delete(k types.Hashable) (deleted bool) {
	n, deleted := delete(t.root, k)
	if deleted {
		// Handling the case of root deletion.
		if t.root.key.Equals(k) {
			t.root = n
		}

		t.count--
	}

	return deleted
}

// delete recursively deletes a key from the tree.
func delete(n *BinaryNode, k types.Hashable) (r *BinaryNode, deleted bool) {
	if r = n; n == nil {
		return nil, false
	}

	if k.Less(n.key) {
		r.left, deleted = delete(n.left, k)
	} else if n.key.Less(k) {
		r.right, deleted = delete(n.right, k)
	} else {
		if n.left != nil && n.right != nil {
			// find the right most element in the left subtree
			s := n.left
			for s.right != nil {
				s = s.right
			}
			r.key = s.key
			r.value = s.value
			r.left, deleted = delete(s, s.key)
		} else if n.left != nil {
			r = n.left
			deleted = true
		} else if n.right != nil {
			r = n.right
			deleted = true
		} else {
			r = nil
			deleted = true
		}
	}

	return
}

// Find returns the value found at the given key.
// Average: O(log(n)) Worst: O(n)
func (t *BinaryTree) Find(k types.Hashable) interface{} {
	return find(t.root, k)
}

func find(n *BinaryNode, k types.Hashable) interface{} {
	if n == nil {
		return nil
	}

	if n.key.Equals(k) {
		return n.value
	} else if k.Less(n.key) {
		return find(n.left, k)
	} else if n.key.Less(k) {
		return find(n.right, k)
	}

	return nil
}

// Clear removes all the nodes from the tree.
// O(n)
func (t *BinaryTree) Clear() {
	t.root = clear(t.root)
	t.count = 0
}

// clear recursively removes all the nodes.
func clear(n *BinaryNode) *BinaryNode {
	if n != nil {
		n.left = clear(n.left)
		n.right = clear(n.right)
	}
	n = nil

	return n
}

// Traverse provides an iterator over the tree.
// O(n)
func (t *BinaryTree) Traverse(tt TraversalType) <-chan interface{} {
	c := make(chan interface{}, t.count)
	go func() {
		switch tt {

		case InOrder:
			inOrder(t.root, c)
		case PreOrder:
			preOrder(t.root, c)
		case PostOrder:
			postOrder(t.root, c)
		}
		close(c)
	}()

	return c
}

// inOrder returns the left, parent, right nodes.
func inOrder(n *BinaryNode, c chan interface{}) {
	if n == nil {
		return
	}

	inOrder(n.left, c)
	c <- n.value
	inOrder(n.right, c)
}

// preOrder returns the parent, left, right nodes.
func preOrder(n *BinaryNode, c chan interface{}) {
	if n == nil {
		return
	}

	c <- n.value
	preOrder(n.left, c)
	preOrder(n.right, c)
}

// postOrder returns the left, right, parent nodes.
func postOrder(n *BinaryNode, c chan interface{}) {
	if n == nil {
		return
	}

	postOrder(n.left, c)
	postOrder(n.right, c)
	c <- n.value
}
