package tree

import ( 
	"fmt"
)

func (node *Node)Traverse() {
	/* if node == nil {
		return
	}
	// 树的遍历，采用中序遍历，先遍历左子树
	node.Left.Traverse()
	node.PrintNode()
	// 遍历右子树
	node.Right.Traverse() */
	node.TraverseFunc(func(node *Node){
		node.PrintNode()
	})
	fmt.Println()
}

func (node *Node) TraverseFunc(f func(*Node)) {
	if node == nil {
		return
	}
	// 树的遍历，采用中序遍历，先遍历左子树
	node.Left.TraverseFunc(f)
	f(node)
	// 遍历右子树
	node.Right.TraverseFunc(f)
}

func (node *Node)TraverseWithChannel() chan *Node{
	out := make(chan *Node)
	go func() {
		node.TraverseFunc(func(node *Node) {
			out <- node
		})
		close(out)
	}()
	return out
}