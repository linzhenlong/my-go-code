package main

import "fmt"

/**
* https://leetcode-cn.com/problems/delete-node-in-a-linked-list/
 */

// ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

func main() {
	node1 := &ListNode{
		Val: 4,
	}
	node2 := &ListNode{
		Val: 5,
	}
	node3 := &ListNode{
		Val: 1,
	}
	node4 := &ListNode{
		Val: 9,
	}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4

	b := node1
	fmt.Println("删除前")
	for b.Next !=nil {
		fmt.Println(b.Val)
		b = b.Next
	}
	fmt.Println(b.Val)

	// 删除第二个节点
	deleteNode(node2)
	a := node1
	fmt.Println("删除后")
	for a.Next !=nil {
		fmt.Println(a.Val)
		a = a.Next
	}
	fmt.Println(a.Val)

}
