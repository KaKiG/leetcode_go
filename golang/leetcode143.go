package main

import "fmt"

func main() {
	var node1, node2, node3, node4, node5 ListNode
	node1.Val = 1
	node2.Val = 2
	node3.Val = 3
	node4.Val = 4
	node5.Val = 5

	node1.Next = &node2
	node2.Next = &node3
	node3.Next = &node4
	node4.Next = &node5

	reorderList(&node1)
	curr := &node1
	for curr != nil {
		fmt.Print(curr.Val)
		curr = curr.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	if head == nil {
		return
	}
	nodes := make([]*ListNode, 0)
	curr := head
	for curr != nil {
		nodes = append(nodes, curr)
		curr = curr.Next
	}

	left := 0
	right := len(nodes) - 1
	for left < right {
		nodes[left].Next = nodes[right]
		left++
		if left >= right {
			break
		}
		nodes[right].Next = nodes[left]
		right--
	}
	nodes[right].Next = nil
}
