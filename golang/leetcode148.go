package main

import "fmt"

func main() {
	var node1, node2, node3, node4, node5 ListNode
	node1.Val = 5
	node2.Val = 4
	node3.Val = 3
	node4.Val = 2
	node5.Val = 6

	node1.Next = &node2
	node2.Next = &node3
	node3.Next = &node4
	node4.Next = &node5

	head := sortList(&node1)
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

func sortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	if head.Next == nil {
		return head
	}

	slow := head
	fast := head
	var prev *ListNode
	for fast != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
		}
	}
	prev.Next = nil

	return merge(sortList(head), sortList(slow))
}

func merge(head1, head2 *ListNode) *ListNode {
	if head1 == nil {
		return head2
	}
	if head2 == nil {
		return head1
	}

	curr1 := head1
	curr2 := head2

	// for curr1 != nil {
	// 	fmt.Print(curr1.Val)
	// 	curr1 = curr1.Next
	// }
	// fmt.Println()
	// for curr2 != nil {
	// 	fmt.Print(curr2.Val)
	// 	curr2 = curr2.Next
	// }
	// fmt.Println()
	// curr1 = list1
	// curr2 = list2

	newHead := &ListNode{Val: 0}
	curr := newHead
	for curr1 != nil && curr2 != nil {
		if curr1.Val < curr2.Val {
			curr.Next = curr1
			curr1 = curr1.Next
		} else {
			curr.Next = curr2
			curr2 = curr2.Next
		}
		curr = curr.Next
	}

	if curr1 == nil {
		curr.Next = curr2
	}
	if curr2 == nil {
		curr.Next = curr1
	}
	return newHead.Next
}
