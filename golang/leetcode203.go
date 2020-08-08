package main

import "fmt"

func main() {
	var node1, node2, node3, node4 ListNode
	node1 = ListNode{2, &node2}
	node2 = ListNode{1, &node3}
	node3 = ListNode{1, &node4}
	node4 = ListNode{2, nil}
	val := 2

	head := removeElements(&node1, val)
	for head != nil {
		fmt.Print(head.Val)
		head = head.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}

	if head.Val == val && head.Next == nil {
		return nil
	}

	// get rid of the node as the head, be careful about nil
	for head != nil {
		if head.Val == val {
			head = head.Next
		} else {
			break
		}
		if head == nil {
			return nil
		}
	}
	slow := head
	fast := slow.Next

	// be careful with nil case
	for fast != nil {
		// get rid of the nodes, if consecutive, throw them away in a loop
		for fast != nil && fast.Val == val {
			slow.Next = fast.Next
			fast = slow.Next
		}
		// move to next
		if fast != nil {
			fast = fast.Next
			slow = slow.Next
		}
	}
	return head
}
