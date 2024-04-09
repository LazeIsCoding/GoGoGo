package Inventory

type Node struct {
	Right *Node
	Left  *Node
	Key   int32
	Index int32
}

type Tree struct {
	Root *Node
}

func New(root *Node) *Tree {
	tr := &Tree{}
	tr.Root = root
	return tr
}

func NewId(key int32, index int32) *Node {
	return &Node{
		Key:   key,
		Index: index,
	}
}

func (tr *Tree) Insert(Item int32, index int32, node *Node) {
	nd := node
	for nd != nil {
		if Item > nd.Key {
			nd = nd.Right
		} else if Item < nd.Key {
			nd = nd.Left
		}
	}

	nd = NewId(Item, index)

}

func (tr *Tree) Search(Item int32) (int32, *Node) {
	nd := tr.Root
	parent := new(Node)
	for nd != nil {
		switch {
		case nd.Key < Item:
			parent = nd
			nd = nd.Right
		case Item < nd.Key:
			parent = nd
			nd = nd.Left
		default:
			return parent.Index, parent
		}
	}
	return -1, parent
}
