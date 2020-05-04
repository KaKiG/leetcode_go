package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	ans := make([]int, 0)
	PartialSolver(root, &ans)
	return ans
}

func PartialSolver(root *TreeNode, ans *[]int) {
	if root.Left != nil {
		PartialSolver(root.Left, ans)
	}
	*ans = append(*ans, root.Val)
	if root.Right != nil {
		PartialSolver(root.Right, ans)
	}
}
