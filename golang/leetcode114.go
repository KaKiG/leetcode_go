package main

import (
	"fmt"
	"strconv"
)

func main() {
	var root, node1, node2, node4, node6, node3 TreeNode
	root = TreeNode{1, &node1, &node2}
	node1 = TreeNode{2, &node3, &node6}
	node2 = TreeNode{5, nil, &node4}
	node3 = TreeNode{3, nil, nil}
	node4 = TreeNode{6, nil, nil}
	//node5 = TreeNode{4, nil, nil}
	node6 = TreeNode{4, nil, nil}
	flatten(&root)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	if root != nil {
		root = lkListTree(preorderTraversal(root))
	}
	fmt.Println(inorderTraversal(root))
}

func lkListTree(nums []int) *TreeNode {
	root := &TreeNode{nums[0], nil, nil}
	if len(nums) == 1 {
		return root
	}
	curr := root
	for i := 1; i < len(nums); i++ {
		curr.Right = &TreeNode{nums[i], nil, nil}
		curr = curr.Right
	}
	fmt.Println(inorderTraversal(root))

	return root
}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	ans := make([]int, 0)
	PartialSolver(root, &ans)
	return ans
}

func PartialSolver(root *TreeNode, ans *[]int) {
	*ans = append(*ans, root.Val)
	if root.Left != nil {
		PartialSolver(root.Left, ans)
	}
	if root.Right != nil {
		PartialSolver(root.Right, ans)
	}
}

func inorderTraversal(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}
	ans := make([]string, 0)
	PartialSolver1(root, &ans)
	return ans
}

func PartialSolver1(root *TreeNode, ans *[]string) {
	if root.Left != nil {
		PartialSolver1(root.Left, ans)
	} else {
		*ans = append(*ans, "L")
	}
	*ans = append(*ans, strconv.Itoa(root.Val))
	if root.Right != nil {
		PartialSolver1(root.Right, ans)
	} else {
		*ans = append(*ans, "R")
	}
}
