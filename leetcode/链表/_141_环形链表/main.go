package main

import "fmt"

type listNode struct {
	Val  int
	Next *listNode
}

func hasCycle(head *listNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	// 快慢指针
	// 让快指针跑的快一点，当fast最终指向了null说明没有环
	// 当fast 与slow相遇了，说明有环，好比跑步比赛里的套圈.
	slow := head
	fast := head.Next
	for fast != nil && fast.Next != nil {
		if slow == fast {
			return true
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return false
}

func main() {

	head := &listNode{
		Val: 3,
	}
	node1 := &listNode{
		Val: 2,
	}
	node2 := &listNode{
		Val: 0,
	}
	node3 := &listNode{
		Val: -4,
	}
	head.Next = node1
	node1.Next = node2
	node2.Next = node3
	//node3.Next = node2
	node3.Next = nil

	fmt.Println(hasCycle(head))

}
