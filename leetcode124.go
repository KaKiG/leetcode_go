package main

import (
	"fmt"
	"math"
)

func main() {
	var root, node1 TreeNode
	root = TreeNode{2, &node1, nil}
	node1 = TreeNode{-1, nil, nil}
	fmt.Println(maxPathSum(&root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	max := math.MinInt32
	solver(root, &max)
	return max
}

func solver(root *TreeNode, max *int) int {
	if root == nil {
		cmd
		return 0
	}
	right := max2(0, solver(root.Right, max))
	left := max2(0, solver(root.Left, max))

	if *max < root.Val+left+right {
		*max = root.Val + left + right
	}
	root.Val += max2(right, left)
	return root.Val
}
func max2(a, b int) int {
	if a > b {
		return a
	}
	return b
}
