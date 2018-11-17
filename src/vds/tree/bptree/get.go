package bptree

import (
	"bytes"
)

import (
	"github.com/timtadh/fs2"
	"github.com/timtadh/fs2/consts"
	"github.com/timtadh/fs2/errors"
)

type bpt_iterator func() (a uint64, idx int, err error, bi bpt_iterator)

func doIter(run func() (fs2.Iterator, error), do func(key, value []byte) error) error {
	kvi, err := run()
	if err != nil {
		return err
	}
	var key, value []byte
	for key, value, err, kvi = kvi(); kvi != nil; key, value, err, kvi = kvi() {
		e := do(key, value)
		if e != nil {
			return e
		}
	}
	return err
}

func doItemIter(run func() (fs2.ItemIterator, error), do func([]byte) error) error {
	it, err := run()
	if err != nil {
		return err
	}
	var item []byte
	for item, err, it = it(); it != nil; item, err, it = it() {
		e := do(item)
		if e != nil {
			return e
		}
	}
	return err
}

// Iterate over all of the key/value pairs in the tree
//
// 	err = bpt.DoIterate(func(key, value []byte) error {
// 		// do something with each key and value in the tree
// 	})
// 	if err != nil {
// 		// handle error
// 	}
//
// Note, it is safe for the keys and values to escape the `do` context.
// They are copied into it so you cannot harm the tree. An unsafe
// version of this is being considered.
func (self *BpTree) DoIterate(do func(key, value []byte) error) error {
	return doIter(
		func() (fs2.Iterator, error) { return self.Iterate() },
		do,
	)
}

// Iterate over each of the keys and values in the tree. I recommend
// that you use the `DoIterate` method instead (it is easier to use). If
// you do use the method always use it as follows:
//
// 	kvi, err := bpt.Iterate()
// 	if err != nil {
// 		// handle error
// 	}
// 	var key, value []byte // must be declared here
// 	// do not use a := assign here only a =
// 	for key, value, err, kvi = kvi(); kvi != nil; key, value, err, kvi = kvi() {
// 		// do something with each key and value
// 	}
// 	// now the iterator could have exited with an error so check the
// 	// error before continuing
// 	if err != nil {
// 		// handle error
// 	}
//
// Note, it is safe for the keys and values to escape the iterator
// context.  They are copied into it so you cannot harm the tree. An
// unsafe version of this is being considered.
func (self *BpTree) Iterate() (kvi fs2.Iterator, err error) {
	return self.Range(nil, nil)
}

// Iterate over all of the keys in the tree. See DoIterate() for usage
// details
func (self *BpTree) DoKeys(do func([]byte) error) error {
	return doItemIter(
		func() (fs2.ItemIterator, error) { return self.Keys() },
		do,
	)
}

// Iterate over all of the keys in the tree. See Iterate() for usage
// details
func (self *BpTree) Keys() (it fs2.ItemIterator, err error) {
	kvi, err := self.Iterate()
	if err != nil {
		return nil, err
	}
	var pk []byte
	it = func() (key []byte, err error, _it fs2.ItemIterator) {
		for key == nil || bytes.Equal(pk, key) {
			key, _, err, kvi = kvi()
			if err != nil {
				return nil, err, nil
			}
			if kvi == nil {
				return nil, nil, nil
			}
		}
		pk = key
		return key, nil, it
	}
	return it, nil
}

// Iterate over all of the values in the tree. See DoIterate() for usage
// details
func (self *BpTree) DoValues(do func([]byte) error) error {
	return doItemIter(
		func() (fs2.ItemIterator, error) { return self.Values() },
		do,
	)
}

// Iterate over all of the values in the tree. See Iterate() for usage
// details
func (self *BpTree) Values() (it fs2.ItemIterator, err error) {
	kvi, err := self.Iterate()
	if err != nil {
		return nil, err
	}
	it = func() (value []byte, err error, _it fs2.ItemIterator) {
		_, value, err, kvi = kvi()
		if err != nil {
			return nil, err, nil
		}
		if kvi == nil {
			return nil, nil, nil
		}
		return value, nil, it
	}
	return it, nil
}

// Iterate over all of the key/values pairs with the given key. See
// DoIterate() for usage details.
func (self *BpTree) DoFind(key []byte, do func(key, value []byte) error) error {
	return doIter(
		func() (fs2.Iterator, error) { return self.Find(key) },
		do,
	)
}

// Iterate over all of the key/values pairs with the given key. See
// Iterate() for usage details.
func (self *BpTree) Find(key []byte) (kvi fs2.Iterator, err error) {
	return self.Range(key, key)
}

// How many key/value pairs are there with the given key.
func (self *BpTree) Count(key []byte) (count int, err error) {
	kvi, err := self.UnsafeRange(key, key)
	if err != nil {
		return 0, err
	}
	count = 0
	for _, _, err, kvi = kvi(); kvi != nil; _, _, err, kvi = kvi() {
		count++
	}
	if err != nil {
		return 0, err
	}
	return count, nil
}

// Iterate over all of the key/values pairs in reverse. See DoIterate()
// for usage details.
func (self *BpTree) DoBackward(do func(key, value []byte) error) error {
	return doIter(
		func() (fs2.Iterator, error) { return self.Backward() },
		do,
	)
}

// Iterate over all of the key/values pairs in reverse. See Iterate()
// for usage details.
func (self *BpTree) Backward() (kvi fs2.Iterator, err error) {
	var bi bpt_iterator
	bi, err = self.backward(nil, nil)
	if err != nil {
		return nil, err
	}
	return self._range(bi)
}

// Iterate over all of the key/values pairs between [from, to]
// inclusive. See DoIterate() for usage details.
func (self *BpTree) DoRange(from, to []byte, do func(key, value []byte) error) error {
	return doIter(
		func() (fs2.Iterator, error) { return self.Range(from, to) },
		do,
	)
}

func (self *BpTree) rangeIterator(from, to []byte) (bi bpt_iterator, err error) {
	if from != nil && to == nil {
		bi, err = self.forward(from, to)
	} else if bytes.Compare(from, to) <= 0 {
		bi, err = self.forward(from, to)
	} else {
		bi, err = self.backward(from, to)
	}
	return bi, err
}

// Iterate over all of the key/values pairs between [from, to]
// inclusive. See Iterate() for usage details.
func (self *BpTree) Range(from, to []byte) (kvi fs2.Iterator, err error) {
	bi, err := self.rangeIterator(from, to)
	if err != nil {
		return nil, err
	}
	return self._range(bi)
}

func (self *BpTree) UnsafeRange(from, to []byte) (kvi fs2.Iterator, err error) {
	bi, err := self.rangeIterator(from, to)
	if err != nil {
		return nil, err
	}
	return self._rangeUnsafe(bi)
}

func (self *BpTree) _range(bi bpt_iterator) (kvi fs2.Iterator, err error) {
	unsafeKvi, err := self._rangeUnsafe(bi)
	if err != nil {
		return nil, err
	}
	// fmt.Println(errors.Errorf("called _range"))
	kvi = func() (key, value []byte, e error, it fs2.Iterator) {
		var k []byte
		var v []byte
		k, v, err, unsafeKvi = unsafeKvi()
		if err != nil {
			return nil, nil, err, nil
		}
		if unsafeKvi == nil {
			return nil, nil, nil, nil
		}
		key = make([]byte, len(k))
		copy(key, k)
		value = make([]byte, len(v))
		// fmt.Println(errors.Errorf("copy"))
		copy(value, v)
		return key, value, nil, kvi
	}
	return kvi, nil
}

func (self *BpTree) _rangeUnsafe(bi bpt_iterator) (kvi fs2.Iterator, err error) {
	kvi = func() (key, value []byte, e error, it fs2.Iterator) {
		var a uint64
		var i int
		a, i, err, bi = bi()
		if err != nil {
			return nil, nil, err, nil
		}
		if bi == nil {
			return nil, nil, nil, nil
		}
		err = self.doKV(a, i, func(k, v []byte) error {
			key = k
			value = v
			return nil
		})
		if err != nil {
			return nil, nil, err, nil
		}
		return key, value, nil, kvi
	}
	return kvi, nil
}

/* returns the key at the address and index or an error
 */
func (self *BpTree) keyAt(a uint64, i int) (key []byte, err error) {
	err = self.do(
		a,
		func(n *internal) error {
			if i >= int(n.meta.keyCount) {
				return errors.Errorf("out of range")
			}
			return n.doKeyAt(self.varchar, i, func(k []byte) error {
				key = make([]byte, len(k))
				copy(key, k)
				return nil
			})
		},
		func(n *leaf) error {
			if i >= int(n.meta.keyCount) {
				return errors.Errorf("out of range")
			}
			return n.doKeyAt(self.varchar, i, func(k []byte) error {
				key = make([]byte, len(k))
				copy(key, k)
				return nil
			})
		},
	)
	if err != nil {
		return nil, err
	}
	return key, nil
}

/* returns the (addr, idx) of the leaf block and the index of the key in
 * the block which has a key greater or equal to the search key.
 */
func (self *BpTree) getStart(key []byte) (a uint64, i int, err error) {
	return self._getStart(self.meta.root, key)
}

func (self *BpTree) _getStart(n uint64, key []byte) (a uint64, i int, err error) {
	var flags consts.Flag
	err = self.bf.Do(n, 1, func(bytes []byte) error {
		flags = consts.AsFlag(bytes)
		return nil
	})
	if err != nil {
		return 0, 0, err
	}
	if flags&consts.INTERNAL != 0 {
		return self.internalGetStart(n, key)
	} else if flags&consts.LEAF != 0 {
		return self.leafGetStart(n, key, false, 0)
	} else {
		return 0, 0, errors.Errorf("Unknown block type")
	}
}

func (self *BpTree) internalGetStart(n uint64, key []byte) (a uint64, i int, err error) {
	var kid uint64
	err = self.doInternal(n, func(n *internal) error {
		if key == nil {
			kid = *n.ptr(0)
			return nil
		}
		i, has, err := find(self.varchar, n, key)
		if err != nil {
			return err
		}
		if !has && i > 0 {
			// if it doesn't have it and the index > 0 then we have the
			// next block so we have to subtract one from the index.
			i--
		}
		kid = *n.ptr(i)
		return nil
	})
	if err != nil {
		return 0, 0, err
	}
	return self._getStart(kid, key)
}

func (self *BpTree) leafGetStart(n uint64, key []byte, stop bool, end uint64) (a uint64, i int, err error) {
	if key == nil {
		return n, 0, nil
	}
	if stop && n == end {
		return 0, 0, errors.Errorf("hit end %v %v %v", n, end, key)
	}
	var next uint64 = 0
	err = self.doLeaf(n, func(n *leaf) (err error) {
		if n.meta.keyCount == 0 {
			// this happens when the tree is empty!
			return nil
		}
		var has bool
		i, has, err = find(self.varchar, n, key)
		if err != nil {
			return err
		}
		if i >= int(n.meta.keyCount) && i > 0 {
			i = int(n.meta.keyCount) - 1
		}
		return n.doKeyAt(self.varchar, i, func(k []byte) error {
			if !has && n.meta.next != 0 && bytes.Compare(k, key) < 0 {
				next = n.meta.next
			}
			return nil
		})
	})
	if err != nil {
		return 0, 0, err
	}
	if next != 0 {
		return self.leafGetStart(next, key, stop, end)
	}
	return n, i, nil
}

func (self *BpTree) lastKey(n uint64) (a uint64, i int, err error) {
	var flags consts.Flag
	err = self.bf.Do(n, 1, func(bytes []byte) error {
		flags = consts.AsFlag(bytes)
		return nil
	})
	if err != nil {
		return 0, 0, err
	}
	if flags&consts.INTERNAL != 0 {
		return self.internalLastKey(n)
	} else if flags&consts.LEAF != 0 {
		return self.leafLastKey(n)
	} else {
		return 0, 0, errors.Errorf("Unknown block type")
	}
}

func (self *BpTree) internalLastKey(n uint64) (a uint64, i int, err error) {
	var kid uint64
	err = self.doInternal(n, func(n *internal) error {
		kid = *n.ptr(int(n.meta.keyCount - 1))
		return nil
	})
	if err != nil {
		return 0, 0, err
	}
	return self.lastKey(kid)
}

func (self *BpTree) leafLastKey(n uint64) (a uint64, i int, err error) {
	var next uint64 = 0
	err = self.doLeaf(n, func(n *leaf) (err error) {
		if n.meta.keyCount == 0 {
			// this happens when the tree is empty!
			return nil
		}
		i = int(n.meta.keyCount) - 1
		return n.doKeyAt(self.varchar, i, func(k []byte) error {
			if n.meta.next != 0 {
				next = n.meta.next
			}
			return nil
		})
	})
	if err != nil {
		return 0, 0, err
	}
	if next != 0 {
		return self.leafLastKey(next)
	}
	return n, i, nil
}

/* returns the (addr, idx) of the leaf block and the index of the key in
 * the block which is either the first key greater than the search key
 * or the last key equal to the search key.
 */
func (self *BpTree) getEnd(key []byte) (a uint64, i int, err error) {
	return self._getEnd(self.meta.root, key)
}

func (self *BpTree) _getEnd(root uint64, key []byte) (a uint64, i int, err error) {
	if key == nil {
		a, i, err = self.lastKey(root)
	} else {
		a, i, err = self._getStart(root, key)
	}
	if err != nil {
		return 0, 0, err
	}
	var equal bool = true
	for equal {
		b, j, end, err := self.nextLoc(a, i)
		if err != nil {
			return 0, 0, err
		}
		if end {
			return a, i, nil
		}
		err = self.doLeaf(b, func(n *leaf) (err error) {
			return n.doKeyAt(self.varchar, j, func(k []byte) error {
				equal = bytes.Equal(key, k)
				return nil
			})
		})
		if err != nil {
			return 0, 0, err
		}
		if equal {
			a, i = b, j
		}
	}
	return a, i, err
}

func (self *BpTree) forward(from, to []byte) (bi bpt_iterator, err error) {
	a, i, err := self.getStart(from)
	if err != nil {
		return nil, err
	} else if from == nil {
		return self.forwardFrom(a, i, to)
	}
	var less bool = false
	err = self.doLeaf(a, func(n *leaf) error {
		if n.meta.keyCount == 0 {
			// this happens when the tree is empty!
			return nil
		}
		return n.doKeyAt(self.varchar, i, func(k []byte) error {
			less = bytes.Compare(k, from) < 0
			return nil
		})
	})
	if err != nil {
		return nil, err
	} else if less {
		bi = func() (uint64, int, error, bpt_iterator) {
			return 0, 0, nil, nil
		}
		return bi, nil
	}
	return self.forwardFrom(a, i, to)
}

func (self *BpTree) forwardFrom(a uint64, i int, to []byte) (bi bpt_iterator, err error) {
	i--
	bi = func() (uint64, int, error, bpt_iterator) {
		var err error
		var end bool
		a, i, end, err = self.nextLoc(a, i)
		if err != nil {
			return 0, 0, err, nil
		}
		if end {
			return 0, 0, nil, nil
		}
		if to == nil {
			return a, i, nil, bi
		}
		var less bool = false
		err = self.doLeaf(a, func(n *leaf) error {
			return n.doKeyAt(self.varchar, i, func(k []byte) error {
				less = bytes.Compare(to, k) < 0
				return nil
			})
		})
		if err != nil {
			return 0, 0, err, nil
		}
		if less {
			return 0, 0, nil, nil
		}
		return a, i, nil, bi
	}
	return bi, nil
}

func (self *BpTree) backward(from, to []byte) (bi bpt_iterator, err error) {
	a, i, err := self.getEnd(from)
	if err != nil {
		return nil, err
	} else if from == nil {
		return self.backwardFrom(a, i, to)
	}
	var greater bool = false
	err = self.doLeaf(a, func(n *leaf) error {
		if n.meta.keyCount == 0 {
			// this happens when the tree is empty!
			return nil
		}
		return n.doKeyAt(self.varchar, i, func(k []byte) error {
			greater = bytes.Compare(k, from) > 0
			return nil
		})
	})
	if err != nil {
		return nil, err
	} else if greater {
		bi = func() (uint64, int, error, bpt_iterator) {
			return 0, 0, nil, nil
		}
		return bi, nil
	}
	return self.backwardFrom(a, i, to)
}

func (self *BpTree) backwardFrom(a uint64, i int, to []byte) (bi bpt_iterator, err error) {
	i++
	bi = func() (uint64, int, error, bpt_iterator) {
		var err error
		var end bool
		a, i, end, err = self.prevLoc(a, i)
		if err != nil {
			return 0, 0, err, nil
		}
		if end {
			return 0, 0, nil, nil
		}
		if to == nil {
			return a, i, nil, bi
		}
		var more bool = false
		err = self.doLeaf(a, func(n *leaf) error {
			return n.doKeyAt(self.varchar, i, func(k []byte) error {
				more = bytes.Compare(to, k) > 0
				return nil
			})
		})
		if err != nil {
			return 0, 0, err, nil
		}
		if more {
			return 0, 0, nil, nil
		}
		return a, i, nil, bi
	}
	return bi, nil
}

func (self *BpTree) nextLoc(a uint64, i int) (uint64, int, bool, error) {
	j := i + 1
	nextBlk := func(a uint64, j int) (uint64, int, bool, error) {
		changed := false
		err := self.doLeaf(a, func(n *leaf) error {
			if j >= int(n.meta.keyCount) && n.meta.next != 0 {
				a = n.meta.next
				j = 0
				changed = true
			}
			return nil
		})
		if err != nil {
			return 0, 0, false, err
		}
		return a, j, changed, nil
	}
	var changed bool = true
	var err error = nil
	for changed {
		a, j, changed, err = nextBlk(a, j)
		if err != nil {
			return 0, 0, false, err
		}
	}
	var end bool = false
	err = self.doLeaf(a, func(n *leaf) error {
		if j >= int(n.meta.keyCount) {
			end = true
		}
		return nil
	})
	if err != nil {
		return 0, 0, false, err
	}
	return a, j, end, nil
}

func (self *BpTree) prevLoc(a uint64, i int) (uint64, int, bool, error) {
	j := i - 1
	prevBlk := func(a uint64, j int) (uint64, int, bool, error) {
		changed := false
		err := self.doLeaf(a, func(n *leaf) error {
			if j < 0 && n.meta.prev != 0 {
				a = n.meta.prev
				changed = true
				return self.doLeaf(a, func(m *leaf) error {
					j = int(m.meta.keyCount) - 1
					return nil
				})
			}
			return nil
		})
		if err != nil {
			return 0, 0, false, err
		}
		return a, j, changed, nil
	}
	var changed bool = true
	var err error = nil
	for changed {
		a, j, changed, err = prevBlk(a, j)
		if err != nil {
			return 0, 0, false, err
		}
	}
	var end bool = false
	err = self.doLeaf(a, func(n *leaf) error {
		if j < 0 || j >= int(n.meta.keyCount) {
			end = true
		}
		return nil
	})
	if err != nil {
		return 0, 0, false, err
	}
	return a, j, end, nil
}
