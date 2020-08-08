/*
 * @lc app=leetcode id=501 lang=golang
 *
 * [501] Find Mode in Binary Search Tree
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
func findMode(root *TreeNode) []int {
	res := make([]int, 0)
	maxCount := 0
	if root == nil {
		return []int{}
	}
	numCountMap := make(map[int]int)
	numCountMap[root.Val]++
	addNum(root.Left, numCountMap)
	addNum(root.Right, numCountMap)

	for k, v := range numCountMap {
		if v > maxCount {
			maxCount = v
			res = []int{k}
		} else if v == maxCount {
			res = append(res, k)
		}
	}
	return res
}

func addNum(root *TreeNode, numCountMap map[int]int) {
	if root == nil {
		return
	}
	numCountMap[root.Val]++
	addNum(root.Left, numCountMap)
	addNum(root.Right, numCountMap)
}

// @lc code=end
