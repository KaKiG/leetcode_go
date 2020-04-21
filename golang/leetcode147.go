package main

import "fmt"

func main() {
	var node1, node2, node3, node4, node5 ListNode
	node1.Val = 5
	node2.Val = 4
	node3.Val = 3
	node4.Val = 2
	node5.Val = 1

	node1.Next = &node2
	node2.Next = &node3
	node3.Next = &node4
	node4.Next = &node5

	head := insertionSortList(&node1)
	curr := head
	for curr != nil {
		fmt.Print(curr.Val)
		curr = curr.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	// list[head...end] is sorted
	end := head
	curr := end.Next

	var prev, after *ListNode
	for curr != nil {
		if curr.Val > end.Val {
			//next node is the last one
			end = curr
			curr = curr.Next
		} else { //insertion
			end.Next = curr.Next
			prev = nil
			after = head

			for after != nil && after.Val < curr.Val {
				prev = after
				after = after.Next
			}

			if prev == nil {
				head = curr
			} else {
				prev.Next = curr
			}
			curr.Next = after
		}
		curr = end.Next
	}
	return head
}
