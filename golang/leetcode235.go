package main

import "fmt"

func main() {
	node1 := TreeNode{1, nil, nil}
	root := TreeNode{2, &node1, nil}

	fmt.Println(lowestCommonAncestor(&root, &root, &node1).Val)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if p.Val <= root.Val && q.Val >= root.Val || p.Val >= root.Val && q.Val <= root.Val {
		return root
	}

	if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}

	return lowestCommonAncestor(root.Right, p, q)
}
