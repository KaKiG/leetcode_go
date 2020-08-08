package main

import "fmt"

func main() {
	m, n := 3, 7
	fmt.Println(uniquePaths(m, n))
}

func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	dp[0][0] = 1
	for i := 1; i < m; i++ {
		dp[i][0] = 1
	}
	for i := 1; i < n; i++ {
		dp[0][i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}
