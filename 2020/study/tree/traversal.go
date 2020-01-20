package tree

func (node *Node)Traverse() {
	if node == nil {
		return
	}
	// 树的遍历，采用中序遍历，先遍历左子树
	node.Left.Traverse()
	node.PrintNode()
	// 遍历右子树
	node.Right.Traverse()
}
