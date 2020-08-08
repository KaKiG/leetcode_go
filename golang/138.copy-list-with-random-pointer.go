package main

/*
 * @lc app=leetcode id=138 lang=golang
 *
 * [138] Copy List with Random Pointer
 */

// @lc code=start

// type Node struct {
// 	Val    int
// 	Next   *Node
// 	Random *Node
// }

// func main() {

// }
func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	newHead := &Node{Val: (*head).Val, Next: nil, Random: nil}

	curr := head
	newCurr := newHead

	nodeMap := make(map[*Node]int)
	nodeArr := make([]*Node, 0)
	i := 0
	for curr != nil {
		nodeMap[curr] = i
		nodeArr = append(nodeArr, newCurr)
		curr = curr.Next
		if curr != nil {
			nextNode := &Node{Val: 0, Next: nil, Random: nil}
			(*nextNode).Val = (*curr).Val
			newCurr.Next = nextNode
			newCurr = newCurr.Next
		}
		i++
	}

	newCurr = newHead
	curr = head
	i = 0
	for curr != nil {
		if curr.Random != nil {
			newCurr.Random = nodeArr[nodeMap[curr.Random]]
		}
		newCurr = newCurr.Next
		curr = curr.Next
		i++
	}
	return newHead

}

// @lc code=end
