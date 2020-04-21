package main

import "fmt"

func main() {
	var root, node1, node2, node4, node6 TreeNode
	root = TreeNode{1, &node1, &node2}
	node1 = TreeNode{2, &node4, &node6}
	node2 = TreeNode{3, nil, nil}
	//node3 = TreeNode{3, nil, nil}
	node4 = TreeNode{4, nil, nil}
	//node5 = TreeNode{4, nil, nil}
	node6 = TreeNode{5, nil, nil}
	fmt.Println(minDepth(&root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if maxDepth(root.Left)-maxDepth(root.Right) == 0 || maxDepth(root.Right)-maxDepth(root.Left) == 1 || maxDepth(root.Right)-maxDepth(root.Left) == -1 {
		return isBalanced(root.Left) && isBalanced(root.Right)
	}
	return false
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	preqNodes := make([]*TreeNode, 0)
	currNodes := preqNodes
	level := 1

	preqNodes = append(preqNodes, root)
	for !(len(preqNodes) == 0 && len(currNodes) == 0) {

		if preqNodes[0].Left != nil {
			currNodes = append(currNodes, preqNodes[0].Left)
		}
		if preqNodes[0].Right != nil {
			currNodes = append(currNodes, preqNodes[0].Right)
		}
		preqNodes = preqNodes[1:]
		if len(preqNodes) == 0 && len(currNodes) != 0 {
			preqNodes = currNodes
			currNodes = currNodes[len(currNodes):]
			level++
		}
	}
	return level
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	preqNodes := make([]*TreeNode, 0)
	currNodes := preqNodes
	level := 1

	preqNodes = append(preqNodes, root)
	for !(len(preqNodes) == 0 && len(currNodes) == 0) {
		if preqNodes[0].Left != nil {
			currNodes = append(currNodes, preqNodes[0].Left)
		}
		if preqNodes[0].Right != nil {
			currNodes = append(currNodes, preqNodes[0].Right)
		}
		if preqNodes[0].Left == nil && preqNodes[0].Right == nil {
			return level
		}
		preqNodes = preqNodes[1:]
		if len(preqNodes) == 0 && len(currNodes) != 0 {
			preqNodes = currNodes
			currNodes = currNodes[len(currNodes):]
			level++
		}
	}
	return level
}
