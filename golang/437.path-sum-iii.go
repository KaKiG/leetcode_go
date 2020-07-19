/*
 * @lc app=leetcode id=437 lang=golang
 *
 * [437] Path Sum III
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, sum int) int {
	res, _ := getSum(root, sum)
	return res
}

func getSum(root *TreeNode, sum int) (int, []int) {
	if root == nil {
		return 0, []int{}
	}
	if isLeaf(root) {
		if sum == root.Val {
			return 1, []int{root.Val}
		} else {
			return 0, []int{root.Val}
		}
	}

	leftSum, leftArr := getSum(root.Left, sum)
	rightSum, rightArr := getSum(root.Right, sum)
	newArr := append(leftArr, rightArr...)
	newArr = append(newArr, 0)

	count := 0
	for i := range newArr {
		newArr[i] += root.Val
		if newArr[i] == sum {
			count++
		}
	}
	return count + leftSum + rightSum, newArr
}
func isLeaf(root *TreeNode) bool {
	if root != nil && root.Left == nil && root.Right == nil {
		return true
	}
	return false
}

// @lc code=end
