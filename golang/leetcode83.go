package main

import "fmt"

func main() {
	var head, n1, n2, n3, n4, n5 ListNode
	var n6, n7 ListNode
	head.Next = &n1
	n1.Next = &n2
	n2.Next = &n3
	n3.Next = &n4
	n4.Next = &n5
	n5.Next = &n6
	n6.Next = &n7
	n7.Next = nil

	n1.Val = 1
	n2.Val = 1
	n3.Val = 1
	n4.Val = 1
	n5.Val = 1
	n6.Val = 1
	n7.Val = 2

	printNode(deleteDuplicates(&n1))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func printNode(head *ListNode) {
	current := head
	for current != nil {
		if current.Next != nil {
			fmt.Print(current.Val, "->")
		} else {
			fmt.Print(current.Val)
			break
		}
		current = current.Next
	}
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	curr := head
	for curr != nil && curr.Next != nil {
		curr = deleteNextDuplicates(curr)
		curr = curr.Next
	}
	return head
}

func deleteNextDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	curr := head
	for curr.Next != nil && curr.Val == curr.Next.Val {
		curr = curr.Next
	}
	head.Next = curr.Next
	return head
}
