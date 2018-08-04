// bptree
package vrdbms

type BNode struct {
	Pageid   int
	Parentid int
	Keys     []int
	KeySize  uint16
	Childs   []int //internal noded points to childs and leafs holds actual page id's
	Leaf     bool
	Modified bool
}

func NewBNode(pageid int, parentid int, leaf bool) *BNode {
	return &BNode{
		Pageid: pageid, Parentid: parentid, Leaf: leaf,
	}
}

type DataNode struct {
}

func (bnode *BNode) insertKey(key int, child int) {
	j := bnode.KeySize
	var nextPos uint16
	if j == 0 {
		nextPos = j
	} else {
		for j >= 0 && bnode.Keys[j] > key {
			bnode.Keys[j+1] = bnode.Keys[j]
			bnode.Childs[j+1] = bnode.Childs[j]
			j--
		}
	}
	bnode.Keys[nextPos] = key
	bnode.Childs[nextPos] = child
	bnode.KeySize++
}

func (bnode *BNode) insertKeyAt(pos int, key int, child int) {
	bnode.Keys[pos] = key
	bnode.Childs[pos] = child
	if bnode.KeySize == 0 {
		bnode.KeySize++
	}

}

func insert(table *Table, key int, child int) {

}

func insertNotFull(node *BNode, key int, child int) {

}

func splitChild(parent *BNode, child1 *BNode, key int, child int) {

}

func splitParent(parent *BNode, child1 *BNode, key int, child int) {

}

func getNode(pageid int) *BNode {
	//	page := Read(pageid)
	return nil
}

func writeNode(node *BNode) {

}

func getParent(node *BNode) *BNode {
	parentid := node.Parentid
	return getNode(parentid)
}
