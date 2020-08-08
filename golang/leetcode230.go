package main

import "fmt"

func main() {
	node5 := TreeNode{1, nil, nil}
	node4 := TreeNode{4, nil, nil}
	node3 := TreeNode{2, &node5, nil}
	node2 := TreeNode{6, nil, nil}
	node1 := TreeNode{3, &node3, &node4}
	root := TreeNode{5, &node1, &node2}
	k := 3
	fmt.Println(kthSmallest(&root, k))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	val := 0
	inorderTraversal(root, &k, &val)
	return val
}

func inorderTraversal(root *TreeNode, k *int, val *int) {

	if root != nil {
		inorderTraversal(root.Left, k, val)
		if *k == 1 {
			*val = root.Val
		}
		*k = *k - 1

		inorderTraversal(root.Right, k, val)
	}
}
