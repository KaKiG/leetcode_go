package main

import "fmt"

func main() {
	node3 := TreeNode{1, nil, nil}
	node2 := TreeNode{20, &node3, nil}
	node1 := TreeNode{9, nil, nil}
	root := TreeNode{3, &node1, &node2}

	fmt.Println(countNodes(&root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	preqNodes := make([]*TreeNode, 0)
	currNodes := preqNodes

	preqNodes = append(preqNodes, root)
	count := 1
	for !(len(preqNodes) == 0 && len(currNodes) == 0) {

		if preqNodes[0].Left != nil {
			currNodes = append(currNodes, preqNodes[0].Left)
			count++
		}
		if preqNodes[0].Right != nil {
			currNodes = append(currNodes, preqNodes[0].Right)
			count++
		}
		preqNodes = preqNodes[1:]
		if len(preqNodes) == 0 && len(currNodes) != 0 {
			preqNodes = currNodes
			currNodes = currNodes[len(currNodes):]
		}
	}
	return count
}
