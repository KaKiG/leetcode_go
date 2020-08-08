package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	var head, num1, num2, num3 ListNode
	head.Next = &num1
	num1.Next = &num2
	num2.Next = &num3
	num3.Next = nil

	num1.Val = 1
	num2.Val = 2
	num3.Val = 3
	k := 2000000000

	fmt.Println(PrintList(rotateRight(head.Next, k)))
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	length := lengthList(head)
	newHead := head
	for i := 0; i < k%length; i++ {
		//fmt.Println(k % lengthList(head))
		newHead = RotateOne(newHead)
	}
	return newHead
}

func lengthList(head *ListNode) int {
	count := 1
	for head.Next != nil {
		head = head.Next
		count++
	}
	return count
}

func RotateOne(head *ListNode) *ListNode {
	tail := head
	for tail.Next.Next != nil {
		tail = tail.Next
	}
	newHead := tail.Next
	tail.Next.Next = head
	tail.Next = nil
	return newHead
}

func PrintList(head *ListNode) []int {
	res := make([]int, 0)
	for head.Next != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	res = append(res, head.Val)
	return res
}
