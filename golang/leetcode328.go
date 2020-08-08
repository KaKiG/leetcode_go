package main

import "fmt"

func main() {
	var node1, node2, node3, node4 ListNode
	node1 = ListNode{1, &node2}
	node2 = ListNode{2, &node3}
	node3 = ListNode{3, &node4}
	node4 = ListNode{4, nil}
	// node5 = ListNode{5, nil}

	curr := oddEvenList(&node1)
	for curr != nil {
		fmt.Print(curr.Val, ", ")
		curr = curr.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return head
	}

	headOddStart, headOddEnd := solver(head.Next.Next)
	headEvenStart := headOddEnd.Next
	headOddEnd.Next = head.Next
	head.Next.Next = headEvenStart
	head.Next = headOddStart

	return head
}

func solver(head *ListNode) (*ListNode, *ListNode) {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return head, head
	}

	headOddStart, headOddEnd := solver(head.Next.Next)
	headEvenStart := headOddEnd.Next
	headOddEnd.Next = head.Next
	head.Next.Next = headEvenStart
	head.Next = headOddStart

	return head, headOddEnd

}
