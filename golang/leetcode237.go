package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(node *ListNode) {
	val := node.Next.Val
	node.Next = node.Next.Next
	node.Val = val
}
