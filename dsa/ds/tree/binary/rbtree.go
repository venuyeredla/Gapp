package binary

import "Gapp/dsa/ds/types"

/*
Properties.
1. Every nodes is either red or black
2. Root is always Black
3. A red node doesn't have red child/parent
3. All nil nodes considered as leafs and need to be blacks
4. Every path from a given node to any of it's descendents nil(leaves) has same no of black nodes.
*/

// While inserting a key we use two tools for balancing tree.
// 1. Recoloring and rotation
func RBPut(n *BinaryNode, key types.Hashable) (root *BinaryNode) {

	if n == nil {
		return &BinaryNode{key: key}
	} else if key.Less(n.key) {
		n.left = RBPut(n.left, key)
		n.left.parent = n

	} else {
		n.right = RBPut(n.right, key)
		n.right.parent = n
	}

	return nil
}

func RBDelete(n *BinaryNode) {

}
