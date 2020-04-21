package main

import (
	"fmt"
	"strconv"
)

func main() {
	var root, node1, node2, node3 TreeNode
	root = TreeNode{1, &node1, &node2}
	node1 = TreeNode{2, nil, &node3}
	node2 = TreeNode{3, nil, nil}
	node3 = TreeNode{5, nil, nil}
	fmt.Println(binaryTreePaths(&root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}

	leftPath := make([]string, 0)
	rightPath := make([]string, 0)

	if root.Left != nil {
		leftPath = binaryTreePaths(root.Left)
	}
	if root.Right != nil {
		rightPath = binaryTreePaths(root.Right)
	}

	if root.Right == nil && root.Left == nil {
		return []string{strconv.Itoa(root.Val)}
	}

	if len(leftPath) != 0 {
		for i := range leftPath {
			leftPath[i] = strconv.Itoa(root.Val) + "->" + leftPath[i]
		}
	}

	if len(rightPath) != 0 {
		for i := range rightPath {
			rightPath[i] = strconv.Itoa(root.Val) + "->" + rightPath[i]
		}
	}

	return append(leftPath, rightPath...)
}
