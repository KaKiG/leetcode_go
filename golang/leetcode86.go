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
	n6.Val = 2
	n7.Val = 2

	printNode(partition(&n1, 3))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	curr := head
	var newList ListNode
	newCurr := &newList
	for curr != nil {
		if curr.Val < x {
			var newNode ListNode
			newNode.Val = curr.Val
			newNode.Next = nil
			newCurr.Next = &newNode
			newCurr = newCurr.Next
		}
		curr = curr.Next
	}
	curr = head
	for curr != nil {
		if curr.Val >= x {
			var newNode ListNode
			newNode.Val = curr.Val
			newNode.Next = nil
			newCurr.Next = &newNode
			newCurr = newCurr.Next
		}
		curr = curr.Next
	}
	return newList.Next
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
