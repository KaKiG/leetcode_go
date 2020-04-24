package main

import "fmt"

func main() {
	prices := []int{1, 4, 2, 7}
	fmt.Println(maxProfit(prices))
}

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	// 0 denotes hold, 1 denotes nothing hold
	dp := make([][]int, len(prices))
	for i := range dp {
		dp[i] = make([]int, 2)
	}
	dp[0][0] = -prices[0]

	for i := 1; i < len(dp); i++ {
		if i > 1 {
			dp[i][0] = max2(dp[i-1][0], dp[i-2][1]-prices[i])
		} else {
			dp[i][0] = max2(dp[i-1][0], dp[i-1][1]-prices[i])
		}

		dp[i][1] = max2(dp[i-1][1], dp[i-1][0]+prices[i])
	}
	return dp[len(prices)-1][1]
	// return dp
}

func max2(a, b int) int {
	if a > b {
		return a
	}
	return b
}
