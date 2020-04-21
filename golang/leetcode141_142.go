package main

import "fmt"

func main() {
	var node1, node2 ListNode
	node1 = ListNode{1, &node2}
	fmt.Println(hasCycle(&node1))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

//----------------------Faster solution-----------------
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	fast := head
	slow := head
	for fast.Next != nil {
		fast = fast.Next
		if fast.Next != nil {
			fast = fast.Next
		} else {
			return false
		}
		slow = slow.Next
		if slow == fast {
			return true
		}
	}
	return false
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return false
	}
	fast := head
	slow := head
	for fast.Next != nil {
		fast = fast.Next
		if fast.Next != nil {
			fast = fast.Next
		} else {
			return false
		}
		slow = slow.Next
		if slow == fast {
			return true
		}
	}
	return false
}

// ---------------------trival solution------------------
// func hasCycle(head *ListNode) bool {
// 	if head == nil {
// 		return false
// 	}
// 	addrMap := make(map[*ListNode]bool)
// 	curr := head
// 	addrMap[head] = true
// 	for curr.Next != nil {
// 		curr = curr.Next
// 		if addrMap[curr] == true {
// 			return true
// 		}
// 		addrMap[curr] = true
// 	}
// 	return false
// }
//
// func detectCycle(head *ListNode) *ListNode {
// 	if head == nil {
// 		return nil
// 	}
// 	addrMap := make(map[*ListNode]bool)
// 	curr := head
// 	addrMap[head] = true
// 	for curr.Next != nil {
// 		curr = curr.Next
// 		if addrMap[curr] == true {
// 			return curr
// 		}
// 		addrMap[curr] = true
// 	}
// 	return nil
// }
