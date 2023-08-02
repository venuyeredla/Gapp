package binary

import (
	"Gapp/dsa/ds/queue"
	"Gapp/dsa/ds/types"
)

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
	queue := new(queue.Queue)
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
