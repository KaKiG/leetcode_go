package main

import "fmt"

func main() {
	n := 10
	fmt.Println(integerBreak(n))
}

func integerBreak(n int) int {
	if n == 0 {
		return 0
	}

	dp := make([]int, n+1)
	dp[1] = 1

	for i := 1; i < n+1; i++ {
		for j := 1; j < i; j++ {
			dp[i] = max2(dp[i], (i-j)*max2(dp[j], j))
		}
	}
	return dp[n]
}

func max2(a, b int) int {
	if a > b {
		return a
	}
	return b
}
