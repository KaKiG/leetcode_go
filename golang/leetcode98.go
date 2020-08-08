package main

import "math"

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//more comfortable way, keep tracking the minimum and maximum value repeatedly.
func isValidBST(root *TreeNode) bool {
	return helper(root, math.MinInt64, math.MaxInt64)
}

func helper(root *TreeNode, low, high int) bool {
	if root == nil {
		return true
	}

	if root.Val <= low || root.Val >= high {
		return false
	}

	return helper(root.Left, low, root.Val) && helper(root.Right, root.Val, high)
}

/*func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if root.Left != nil {
		if root.Left.Val >= root.Val {
			return false
		}
	}
	if root.Right != nil {
		if root.Right.Val <= root.Val {
			return false
		}
	}

	return isValidBST(root.Left) && isValidBST(root.Right) && (maxT(root.Left) < root.Val || root.Left == nil) && (minT(root.Right) > root.Val || root.Right == nil)
}

func maxT(root *TreeNode) int {
	if root == nil {
		return math.MinInt64
	}
	return max2(max2(maxT(root.Left), root.Val), maxT(root.Right))
}

func minT(root *TreeNode) int {
	if root == nil {
		return math.MaxInt64
	}
	return min2(min2(minT(root.Left), root.Val), minT(root.Right))
}

func min2(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max2(a, b int) int {
	if a > b {
		return a
	}
	return b
}
*/
