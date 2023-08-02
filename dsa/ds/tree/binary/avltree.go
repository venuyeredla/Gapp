package binary

import (
	"Gapp/dsa/ds/errors"
	"Gapp/dsa/ds/types"
)

func (self *BinaryNode) AvlPut(key types.Hashable, value interface{}) (_ *BinaryNode, updated bool) {
	if self == nil {
		return &BinaryNode{key: key, value: value, height: 1}, false
	}

	if self.key.Equals(key) {
		self.value = value
		return self, true
	}

	if key.Less(self.key) {
		self.left, updated = self.left.AvlPut(key, value)
	} else {
		self.right, updated = self.right.AvlPut(key, value)
	}
	if !updated {
		self.height += 1
		return self.balance(), updated
	}
	return self, updated
}

func (self *BinaryNode) AvlRemove(key types.Hashable) (_ *BinaryNode, value interface{}, err error) {
	if self == nil {
		return nil, nil, errors.NotFound(key)
	}

	if self.key.Equals(key) {
		if self.left != nil && self.right != nil {
			if self.left.Size() < self.right.Size() {
				lmd := self.right.lmd()
				lmd.left = self.left
				return self.right, self.value, nil
			} else {
				rmd := self.left.rmd()
				rmd.right = self.right
				return self.left, self.value, nil
			}
		} else if self.left == nil {
			return self.right, self.value, nil
		} else if self.right == nil {
			return self.left, self.value, nil
		} else {
			return nil, self.value, nil
		}
	}
	if key.Less(self.key) {
		self.left, value, err = self.left.AvlRemove(key)
	} else {
		self.right, value, err = self.right.AvlRemove(key)
	}
	if err != nil {
		return self.balance(), value, err
	}
	return self, value, err
}

func (self *BinaryNode) pop_node(node *BinaryNode) *BinaryNode {
	if node == nil {
		panic("node can't be nil")
	} else if node.left != nil && node.right != nil {
		panic("node must not have both left and right")
	}

	if self == nil {
		return nil
	} else if self == node {
		var n *BinaryNode
		if node.left != nil {
			n = node.left
		} else if node.right != nil {
			n = node.right
		} else {
			n = nil
		}
		node.left = nil
		node.right = nil
		return n
	}

	if node.key.Less(self.key) {
		self.left = self.left.pop_node(node)
	} else {
		self.right = self.right.pop_node(node)
	}

	self.height = max(self.left.Height(), self.right.Height()) + 1
	return self
}

func (self *BinaryNode) push_node(node *BinaryNode) *BinaryNode {
	if node == nil {
		panic("node can't be nil")
	} else if node.left != nil || node.right != nil {
		panic("node now be a leaf")
	}

	if self == nil {
		node.height = 1
		return node
	} else if node.key.Less(self.key) {
		self.left = self.left.push_node(node)
	} else {
		self.right = self.right.push_node(node)
	}
	self.height = max(self.left.Height(), self.right.Height()) + 1
	return self
}

func (self *BinaryNode) rotate_right() *BinaryNode {
	if self == nil {
		return self
	}
	if self.left == nil {
		return self
	}
	new_root := self.left.rmd()
	self = self.pop_node(new_root)
	new_root.left = self.left
	new_root.right = self.right
	self.left = nil
	self.right = nil
	return new_root.push_node(self)
}

func (self *BinaryNode) rotate_left() *BinaryNode {
	if self == nil {
		return self
	}
	if self.right == nil {
		return self
	}
	new_root := self.right.lmd()
	self = self.pop_node(new_root)
	new_root.left = self.left
	new_root.right = self.right
	self.left = nil
	self.right = nil
	return new_root.push_node(self)
}

func (self *BinaryNode) balance() *BinaryNode {
	if self == nil {
		return self
	}
	for abs(self.left.Height()-self.right.Height()) > 2 {
		if self.left.Height() > self.right.Height() {
			self = self.rotate_right()
		} else {
			self = self.rotate_left()
		}
	}
	return self
}

func (self *BinaryNode) Height() int {
	if self == nil {
		return 0
	}
	return self.height
}

func (self *BinaryNode) Size() int {
	if self == nil {
		return 0
	}
	return 1 + self.left.Size() + self.right.Size()
}

func (self *BinaryNode) Left() types.BinaryTreeNode {
	if self.left == nil {
		return nil
	}
	return self.left
}

func (self *BinaryNode) Right() types.BinaryTreeNode {
	if self.right == nil {
		return nil
	}
	return self.right
}

func (self *BinaryNode) _md(side func(*BinaryNode) *BinaryNode) *BinaryNode {
	if self == nil {
		return nil
	} else if side(self) != nil {
		return side(self)._md(side)
	} else {
		return self
	}
}

func (self *BinaryNode) lmd() *BinaryNode {
	return self._md(func(node *BinaryNode) *BinaryNode { return node.left })
}

func (self *BinaryNode) rmd() *BinaryNode {
	return self._md(func(node *BinaryNode) *BinaryNode { return node.right })
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
