package main

import (
	"fmt"
	"math"
)

func main() {
	coins := []int{1, 2, 5}
	amount := 11
	fmt.Println(coinChange(coins, amount))
}

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)

	for i := range dp {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0
	for i := 1; i < amount+1; i++ {
		for j := range coins {
			if i >= coins[j] {
				dp[i] = min2(dp[i], dp[i-coins[j]]+1)
			}
		}
	}
	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}

func min2(a, b int) int {
	if b < a {
		return b
	}
	return a
}
