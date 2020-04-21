package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	if sum == root.Val && root.Left == nil && root.Right == nil {
		return true
	}
	return hasPathSum(root.Left, sum-root.Val) || hasPathSum(root.Right, sum-root.Val)
}

func pathSum(root *TreeNode, sum int) [][]int {
	if root == nil {
		return [][]int{}
	}
	ans := make([][]int, 0)
	Solver(root, sum, nil, &ans)
	return ans
}

func Solver(root *TreeNode, sum int, curr []int, ans *[][]int) {

	if root.Val == sum && root.Left == nil && root.Right == nil {
		curr = append(curr, root.Val)
		*ans = append(*ans, append([]int{}, curr...))
	}

	if root.Left != nil {
		Solver(root.Left, sum-root.Val, append(curr, root.Val), ans)
	}
	if root.Right != nil {
		Solver(root.Right, sum-root.Val, append(curr, root.Val), ans)
	}
}
