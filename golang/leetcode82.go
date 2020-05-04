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
	n4.Val = 2
	n5.Val = 3
	n6.Val = 5
	n7.Val = 6

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
	var newHead ListNode
	newHead.Next = head
	curr := &newHead
	for curr.Next != nil {
		if curr.Next.Next == nil {
			return newHead.Next
		}
		if curr.Next.Val != curr.Next.Next.Val {
			curr = curr.Next
		}
		curr = deleteNextDuplicates(curr)
	}
	return newHead.Next
}

func deleteNextDuplicates(head *ListNode) *ListNode {
	if head.Next == nil || head.Next.Next == nil {
		return head
	}
	check := head
	tmp := 0
	if check.Next.Val == check.Next.Next.Val {
		tmp = check.Next.Val
	} else {
		return head
	}
	for check.Next != nil && check.Next.Val == tmp {
		check = check.Next
	}
	head.Next = check.Next
	return head
}
