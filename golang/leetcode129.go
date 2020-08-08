package main

import (
	"fmt"
	"math"
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
	fmt.Println(sumNumbers(&root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	ans := make([][]int, 0)
	dfs(root, nil, &ans)
	return Cal2(ans)
}

func Cal2(nums [][]int) int {
	sum := 0
	for i := range nums {
		sum += Cal1(nums[i])
	}
	return sum
}

func Cal1(nums []int) int {
	n := len(nums)
	sum := 0
	for i := range nums {
		sum += nums[i] * int(math.Pow10(n-i-1))
	}
	return sum
}

func dfs(root *TreeNode, curr []int, ans *[][]int) {
	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil {
		*ans = append(*ans, append([]int{}, append(curr, root.Val)...))
		return
	}

	dfs(root.Left, append(curr, root.Val), ans)
	dfs(root.Right, append(curr, root.Val), ans)
}
