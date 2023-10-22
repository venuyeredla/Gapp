package binary

import (
	"Gapp/dsa/ds/errors"
	"Gapp/dsa/ds/stackqueue"
	"Gapp/dsa/ds/tree"
	"Gapp/dsa/ds/types"
)

type COLOR uint
type TreeType uint

// TraversalType represents one of the three know traversals.
type TraversalType int

const (
	InOrder TraversalType = iota
	PreOrder
	PostOrder
)

const (
	AVL TreeType = iota
	RB
	LEVEL
	BST
)

const (
	RED COLOR = iota
	BLACK
)

type BinaryNode struct {
	key                 types.Hashable
	value               interface{}
	parent, left, right *BinaryNode
	height              int   //For avl tree
	color               COLOR // For redblack tree
}

type BinaryTree struct {
	root     *BinaryNode
	treeType TreeType
	count    int
}

func NewBinaryTree(tType TreeType) *BinaryTree {
	return &BinaryTree{treeType: tType}
}

func NewAvlTree() *BinaryTree {
	return &BinaryTree{treeType: AVL}
}

func (self *BinaryTree) Iterate() types.KVIterator {
	return self.root.Iterate()
}

func (self *BinaryTree) Items() (vi types.KIterator) {
	return types.MakeItemsIterator(self)
}

func (self *BinaryTree) Values() types.Iterator {
	return self.root.Values()
}

func (self *BinaryTree) Keys() types.KIterator {
	return self.root.Keys()
}

func (self *BinaryTree) Root() types.TreeNode {
	return self.root
}

func (self *BinaryTree) Size() int {
	return self.root.Size()
}

func (self *BinaryTree) Has(key types.Hashable) bool {
	return self.root.Has(key)
}

func (self *BinaryNode) Has(key types.Hashable) (has bool) {
	if self == nil {
		return false
	}
	if self.key.Equals(key) {
		return true
	} else if key.Less(self.key) {
		return self.left.Has(key)
	} else {
		return self.right.Has(key)
	}
}

func (self *BinaryTree) Get(key types.Hashable) (value interface{}, err error) {
	return self.root.Get(key)
}

func (self *BinaryNode) Get(key types.Hashable) (value interface{}, err error) {
	if self == nil {
		return nil, errors.NotFound(key)
	}
	if self.key.Equals(key) {
		return self.value, nil
	} else if key.Less(self.key) {
		return self.left.Get(key)
	} else {
		return self.right.Get(key)
	}
}

// TreeNode interface
func (self *BinaryNode) Key() types.Hashable {
	return self.key
}

func (self *BinaryNode) Value() interface{} {
	return self.value
}

func (self *BinaryNode) GetChild(i int) types.TreeNode {
	return types.DoGetChild(self, i)
}

func (self *BinaryNode) ChildCount() int {
	return types.DoChildCount(self)
}

func (self *BinaryNode) Children() types.TreeNodeIterator {
	return types.MakeChildrenIterator(self)
}

// TreeNode interface

func (self *BinaryNode) Iterate() types.KVIterator {
	tni := tree.TraverseBinaryTreeInOrder(self)
	return types.MakeKVIteratorFromTreeNodeIterator(tni)
}

func (self *BinaryNode) Keys() types.KIterator {
	return types.MakeKeysIterator(self)
}

func (self *BinaryNode) Values() types.Iterator {
	return types.MakeValuesIterator(self)
}

func (self *BinaryTree) Put(key types.Hashable, value interface{}) (err error) {
	switch self.treeType {
	case RB:
		self.root = RBPut(self.root, key)
		break
	case AVL:
		self.root, _ = self.root.AvlPut(key, value)
		return nil
	case LEVEL:
		self.root, _ = self.root.LevelPut(key, value) //// Average: O(log(n)) Worst: O(n)
		return nil
	case BST:
		self.root, _ = self.root.BstPut(key, value) //// Average: O(log(n)) Worst: O(n)
		return nil
	}
	return nil
}

func (self *BinaryTree) Remove(key types.Hashable) (value interface{}, err error) {

	switch self.treeType {
	case RB:
		self.root = RBPut(self.root, key)
		break
	case AVL:
		new_root, value, err := self.root.AvlRemove(key)
		if err != nil {
			return nil, err
		}
		self.root = new_root
		return value, nil
	case LEVEL:
		RBPut(self.root, key)
		break
	}

	return value, nil
}

// insert recusively adds a key+value in the tree.
// Lever order insertion using Queue
func (self *BinaryNode) LevelPut(key types.Hashable, value interface{}) (r *BinaryNode, updated bool) {

	//func insertInOrder(n *node, k int, v interface{}) (r *node, added bool) {
	newNode := &BinaryNode{key: key, value: value}
	if r = self; self == nil {
		r = newNode
		updated = true
		return r, updated
	}
	queue := new(stackqueue.Queue)
	queue.Init()
	queue.Push(self)

	for !queue.IsEmpty() {
		current := ToNode(queue.Pop())
		if current.left == nil {
			current.left = newNode
			break
		} else {
			queue.Push(current.left)
		}
		if current.right == nil {
			current.right = newNode
			break
		} else {
			queue.Push(current.right)
		}
	}

	return r, true
}

func ToNode(n1 interface{}) (n *BinaryNode) {
	//node(n)
	return
}
