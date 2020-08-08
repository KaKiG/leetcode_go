package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	len1, len2 := 0, 0
	curr1 := headA
	curr2 := headB

	for curr1 != nil {
		len1++
		curr1 = curr1.Next
	}

	for curr2 != nil {
		len2++
		curr2 = curr2.Next
	}
	var long, short *ListNode
	var diff int
	if len1 > len2 {
		long, short, diff = headA, headB, len1-len2
	} else {
		long, short, diff = headB, headA, len2-len1
	}

	for i := 0; i < diff; i++ {
		long = long.Next
	}

	for {
		if long == short {
			return long
		}
		long = long.Next
		short = short.Next
	}
	return nil
}

// func getIntersectionNode(headA, headB *ListNode) *ListNode {
// 	curr1 := headA
// 	curr2 := headB
// 	list := make(map[*ListNode]bool)
//
// 	for curr1 != nil {
// 		list[curr1] = true
// 		curr1 = curr1.Next
// 	}
//
// 	for curr2 != nil {
// 		if list[curr2] {
// 			return curr2
// 		}
// 		curr2 = curr2.Next
// 	}
// 	return nil
// }
