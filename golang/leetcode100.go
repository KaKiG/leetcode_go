package main

import (
	"fmt"
	"strconv"
)

func main() {
	var root1, left1, root2, right2 TreeNode
	root1.Left = &left1
	root1.Right = nil
	root1.Val = 0
	left1.Left = nil
	left1.Right = nil
	left1.Val = -5

	root2.Left = &right2
	root2.Right = nil
	root2.Val = 0
	right2.Left = nil
	right2.Right = nil
	right2.Val = -8

	fmt.Println(isSameTree(&root1, &root2))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	ans1 := inorderTraversal(p)
	ans2 := inorderTraversal(q)

	if len(ans1) != len(ans2) {
		return false
	}
	fmt.Println(ans1, ans2)
	for i := range ans1 {
		if ans1[i] != ans2[i] {
			return false
		}
	}
	return true
}

func inorderTraversal(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}
	ans := make([]string, 0)
	PartialSolver(root, &ans)
	return ans
}

func PartialSolver(root *TreeNode, ans *[]string) {
	if root.Left != nil {
		PartialSolver(root.Left, ans)
	} else {
		*ans = append(*ans, "l")
	}
	*ans = append(*ans, strconv.Itoa(root.Val))
	if root.Right != nil {
		PartialSolver(root.Right, ans)
	} else {
		*ans = append(*ans, "r")
	}
}
