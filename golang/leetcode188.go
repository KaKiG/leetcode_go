package main

import (
	"fmt"
	"math"
)

func main() {
	prices := []int{3, 2, 6, 5, 0, 3}
	k := 2
	fmt.Print(maxProfit(k, prices))
}

func maxProfit(k int, prices []int) int {
	// when k is greater than len(prices) / 2
	// we degenerate it into easier way
	if k > len(prices)/2 {
		return norMax(prices)
	}
	dp := make([][][]int, len(prices)+1)
	for i := range dp {
		dp[i] = make([][]int, k+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, 2)
		}
	}

	// make sure, you cannot gain profit via selling stock in day 1
	for j := 0; j < k+1; j++ {
		dp[0][j][1] = math.MinInt32
	}

	// here we regard it as max profit on day i, at most j transactions, whether on hold stock
	// we can define whether selling or buying a stock is a transaction
	// however, whatever we define, we will deal with the transaction completion first (i.e. the recursion term with j-1)
	// otherwise it will result in sell a stock you did not buy yet
	for i := range prices {
		for j := 1; j < k+1; j++ {
			dp[i+1][j][1] = max2(dp[i][j-1][0]-prices[i], dp[i][j][1])
			dp[i+1][j][0] = max2(dp[i][j][1]+prices[i], dp[i][j][0])
		}
	}
	return dp[len(prices)][k][0]
}

func max2(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func norMax(prices []int) int {
	maxprofit := 0
	for i := 1; i < len(prices); i++ {
		maxprofit += max2(0, prices[i]-prices[i-1])
	}
	return maxprofit
}
