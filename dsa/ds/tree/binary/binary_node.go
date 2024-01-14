package binary

import "Gapp/dsa/ds/types"

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
	LEVEL TreeType = iota
	AVL
	BST
	RB
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
