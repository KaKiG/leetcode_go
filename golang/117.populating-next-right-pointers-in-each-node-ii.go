/*
 * @lc app=leetcode id=117 lang=golang
 *
 * [117] Populating Next Right Pointers in Each Node II
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	if root.Right != nil {
		root.Right.Next = findNext(root)
	}
	if root.Left != nil {
		if root.Right != nil {
			root.Left.Next = root.Right
		} else {
			root.Left.Next = findNext(root)
		}
	}

	// @note, root.Right should go first
	connect(root.Right)
	connect(root.Left)
	return root
}

func findNext(root *Node) *Node {
	curr := root.Next
	for curr != nil {
		if curr.Left != nil {
			return curr.Left
		}
		if curr.Right != nil {
			return curr.Right
		}
		curr = curr.Next
	}
	return nil
}

// @lc code=end
