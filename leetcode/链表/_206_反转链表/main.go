package main

import "fmt"

/**
https://leetcode-cn.com/problems/reverse-linked-list/
输入: 1->2->3->4->5->NULL
输出: 5->4->3->2->1->NULL
**/

// ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}
//PrintListNode ...
func (l *ListNode)PrintListNode() {
	for l.Next != nil {
		fmt.Println(l.Val)
		l = l.Next
	}
	fmt.Println(l.Val)
}

// 递归的方式
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	newNode := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newNode
}

func main() {
	head := &ListNode{
		Val: 1,
	}
	node2 := &ListNode{
		Val: 2,
	}
	head.Next = node2
	node3 := &ListNode{
		Val: 3,
	}
	node2.Next = node3
	node4 := &ListNode{
		Val: 4,
	}
	node3.Next = node4
	node5 := &ListNode{
		Val: 5,
	}
	node4.Next = node5
	head.PrintListNode()
	newHead := reverseList(head)
	fmt.Println("=========")
	newHead.PrintListNode()
}
