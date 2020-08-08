package main

import "fmt"

func main() {
	var root, node1, node2, node3, node4 TreeNode
	root = TreeNode{1, &node1, nil}
	node1 = TreeNode{3, nil, &node2}
	node2 = TreeNode{2, nil, nil}
	node3 = TreeNode{2, nil, nil}
	node4 = TreeNode{5, nil, nil}
	ansInt, _ := inorderTraversal(&root)
	fmt.Println(ansInt)
	fmt.Println(findSwap(ansInt))
	recoverTree(&root)
	ansInt2, _ := inorderTraversal(&root)
	fmt.Println(ansInt2)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recoverTree(root *TreeNode) {
	ints, nodes := inorderTraversal(root)
	p1, p2 := findSwap(ints)
	nodes[p1].Val, nodes[p2].Val = nodes[p2].Val, nodes[p1].Val
}

func findSwap(nums []int) (int, int) {
	start, end := -1, -1
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1] < nums[i] && start == -1 {
			start = i
		}
		if nums[i+1] < nums[i] {
			end = i + 1
		}
	}
	return start, end
}

func inorderTraversal(root *TreeNode) ([]int, []*TreeNode) {
	if root == nil {
		return []int{}, []*TreeNode{}
	}
	ansInt := make([]int, 0)
	ansNode := make([]*TreeNode, 0)
	PartialSolverT(root, &ansInt, &ansNode)
	return ansInt, ansNode
}

func PartialSolverT(root *TreeNode, ansInt *[]int, ansNode *[]*TreeNode) {
	if root.Left != nil {
		PartialSolverT(root.Left, ansInt, ansNode)
	}
	*ansInt = append(*ansInt, root.Val)
	*ansNode = append(*ansNode, root)
	if root.Right != nil {
		PartialSolverT(root.Right, ansInt, ansNode)
	}
}
