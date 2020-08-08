import (
	"math"
)

/*
 * @lc app=leetcode id=375 lang=golang
 *
 * [375] Guess Number Higher or Lower II
 */

// @lc code=start
func getMoneyAmount(n int) int {
	dp := make([][]int, n+1) //dp[row, col] means min cost [row, col]
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for col := 1; col <= n; col++ {
		for row := col; row >= 1; row-- {
			if row == col {
				continue
			}

			if row == col-1 {
				dp[row][col] = row
				continue
			}
			minCost := math.MaxInt32
			for j := row + 1; j <= col-1; j++ {
				minCost = min2(minCost, j+max2(dp[row][j-1], dp[j+1][col]))
			}
			dp[row][col] = minCost
		}
	}
	return dp[1][n]
}

func max2(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min2(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
