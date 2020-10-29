package main

import (
	"fmt"

	"github.com/linzhenlong/my-go-code/leetcode/链表/listnode"
)

/**
 https://leetcode-cn.com/problems/remove-linked-list-elements/
**/

func removeElements(head *listnode.ListNode, val int) *listnode.ListNode {
	if head == nil {
		return head
	}
	tmp := head // 临时变量遍历使用
	for tmp.Next != nil {
		// 判断当前训练的节点的下一个节点的val是否与被删除的val 相等
		if tmp.Next.Val == val {
			// 如果相等那么将当前的循环的next是要被一层的节点，所以将当前的next 指向当前的下一个的下一个
			tmp.Next = tmp.Next.Next
		} else {
			// 如果不相等，正常遍历，挪动.
			tmp = tmp.Next
		}
	}
	// 由于上面的while循环是判断当前的下一个，这里还要判断当前的是否需要山删除
	if head.Val == val {
		return head.Next
	} 
	return head
}
// 辅助头节点的方式删除
func removeElements2(head *listnode.ListNode, val int) *listnode.ListNode {
	if head == nil {
		return head
	}
	helperHeadNode := &listnode.ListNode{
		Val: 0,
		Next: head,
	}
	tmp := helperHeadNode
	for tmp.Next !=nil {
		if tmp.Next.Val == val {
			tmp.Next = tmp.Next.Next
		} else {
			tmp = tmp.Next
		}
	}
	return helperHeadNode.Next
}

func main() {
	node1 := &listnode.ListNode{
		Val: 1,
	}
	node2 := &listnode.ListNode{
		Val: 1,
	}
	node3 := &listnode.ListNode{
		Val: 1,
	}
	node4 := &listnode.ListNode{
		Val: 3,
	}
	node5 := &listnode.ListNode{
		Val: 4,
	}
	node6 := &listnode.ListNode{
		Val: 5,
	}
	node7 := &listnode.ListNode{
		Val: 6,
	}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	node5.Next = node6
	node6.Next = node7

	
	b := node1
	for b != nil{
		fmt.Println("删除前:",b.Val)
		b = b.Next
	}
	c := node1
	removeElements(c,1)
	fmt.Println("分割线=========")
	for c != nil{
		fmt.Println("删除后:",c.Val)
		c = c.Next
	}
}
