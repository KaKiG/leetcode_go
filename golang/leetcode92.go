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
	n2.Val = 2
	n3.Val = 3
	n4.Val = 4
	n5.Val = 5
	n6.Val = 6
	n7.Val = 7

	m := 1
	n := 1
	printNode(reverseBetween(&n1, m, n))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head.Next == nil {
		return head
	}
	if m == 1 {
		return reverseFirstK(head, n).Next
	}
	curr := head
	link := head
	for i := 0; i < m-1; i++ {
		curr = curr.Next
		if i == 0 {
			continue
		}
		link = link.Next
	}
	curr = reverseFirstK(curr, n-m+1)
	link.Next = curr.Next
	return head
}

func reverseFirstK(head *ListNode, k int) *ListNode {
	pHead := &ListNode{0, head}
	current := pHead
	for i := 1; i <= k; i++ {
		current = pHead
		for j := 1; j < i; j++ {
			current = current.Next
		}
		if current.Next != nil {
			pHead = swapTwoNodes(pHead, current)
		}
	}
	return pHead
}

func swapTwoNodes(pFirst, pSecond *ListNode) *ListNode {
	toNode2 := pSecond.Next

	if toNode2.Next != nil {
		pSecond.Next = toNode2.Next
	} else {
		pSecond.Next = nil
	}

	toNode2.Next = pFirst.Next
	pFirst.Next = toNode2

	return pFirst
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
