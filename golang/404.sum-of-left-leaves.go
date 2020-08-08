/*
 * @lc app=leetcode id=404 lang=golang
 *
 * [404] Sum of Left Leaves
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}

	sum := 0
	if isLeave(root.Left) {
		sum += root.Left.Val
	} else {
		sum += sumOfLeftLeaves(root.Left)
	}
	sum += sumOfLeftLeaves(root.Right)
	return sum
}

func isLeave(root *TreeNode) bool {
	if root != nil && root.Left == nil && root.Right == nil {
		return true
	}
	return false
}

// @lc code=end
