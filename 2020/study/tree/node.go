package tree

import "fmt"

type Node struct {
	Left, Right *Node
	Value int
}
func CreateNode(value int) *Node {
	return &Node{Value: value}
}
func (node *Node)PrintNode() {
	fmt.Println(node.Value, " ")
}
func (node *Node)SetValue(v int) {
	if node == nil {
		fmt.Println("setting value to nil node")
		return
	}
	node.Value = v
}
