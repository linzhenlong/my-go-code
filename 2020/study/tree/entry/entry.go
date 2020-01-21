package main

import (
	"github.com/linzhenlong/my-go-code/2020/study/tree"
	"golang.org/x/tools/container/intsets"
	"fmt"
)

type myTreeNode struct {
	node *tree.Node
}
func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil{
		return
	}
	left := myTreeNode{myNode.node.Left}
	left.postOrder()
	right := myTreeNode{myNode.node.Right}
	right.postOrder()
	myNode.node.PrintNode()

}
func testSparse() {
	s := intsets.Sparse{}
	s.Insert(1)
	s.Insert(1000)
	s.Insert(1000000)
	fmt.Println(s.Has(1000))
	fmt.Println(s.Has(100000))
}

func main() {

	var root tree.Node
	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{nil,nil,5}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)
	root.Right.Left.SetValue(4)
	root.Traverse()
	fmt.Println()
	myRoot := myTreeNode{
		node: &root,
	}
	myRoot.postOrder()
	
	testSparse()

	nodeCount := 0
	root.TraverseFunc(
		func(node *tree.Node) {
			nodeCount ++
		})
	fmt.Println("node count:", nodeCount)
}