package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	return max2(solver(root))
}

// {not rob this level, rob this level}
func solver(root *TreeNode) [2]int {
	if root == nil {
		return [2]int{0, 0}
	}

	left := solver(root.Left)
	right := solver(root.Right)

	return [2]int{max2(left) + max2(right), root.Val + left[0] + right[0]}
}

func max2(num [2]int) int {
	if num[0] > num[1] {
		return num[0]
	}
	return num[1]
}
