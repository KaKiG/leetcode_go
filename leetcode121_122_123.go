package main

import "fmt"

func main() {
	prices := []int{1, 2, 3, 4, 5}
	fmt.Println(maxProfits(prices))
}

func maxProfitss(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	dp := make([]int, len(prices))
	for i := 1; i < len(dp); i++ {
		for j := 0; j <= i; j++ {
			if j == 0 {
				dp[i] = max2(maxProfit(prices[:i+1]), dp[i])
				continue
			}
			dp[i] = max2(maxProfit(prices[:j])+maxProfit(prices[j:i+1]), dp[i])
		}
	}
	return maxArray(dp)
}

func maxProfits(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	res := 0
	for i := 1; i < len(prices); i++ {
		res += max2(prices[i]-prices[i-1], 0)
	}
	return res
}

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	for i := range prices {
		if i == len(prices)-1 {
			return 0
		}
		if prices[i+1] > prices[i] {
			prices = prices[i:]
			break
		}
	}
	dp := make([]int, len(prices))

	for i := 1; i < len(dp); i++ {
		dp[i] = max2(prices[i]-prices[i-1]+dp[i-1], 0)
	}
	return maxArray(dp)
}

func max2(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxArray(nums []int) int {
	max := nums[0]
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	return max
}
