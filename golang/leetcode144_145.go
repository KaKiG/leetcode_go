package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	ans := make([]int, 0)
	preOrderSolver(root, &ans)
	return ans
}

func preOrderSolver(root *TreeNode, ans *[]int) {
	if root != nil {
		*ans = append(*ans, root.Val)
		preOrderSolver(root.Left, ans)
		preOrderSolver(root.Right, ans)
	}
}

func postorderTraversal(root *TreeNode) []int {
	ans := make([]int, 0)
	postOrderSolver(root, &ans)
	return ans
}

func postOrderSolver(root *TreeNode, ans *[]int) {
	if root != nil {
		postOrderSolver(root.Left, ans)
		postOrderSolver(root.Right, ans)
		*ans = append(*ans, root.Val)
	}
}
