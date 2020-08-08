package main

import "fmt"

func main() {
	nums := []int{3, 1, 5, 8}
	fmt.Println(maxCoins(nums))
}

// dynamic programming, we split nums in to different intervals
// opt(i,j) := the maximum value at interval nums[i,j+1]
func maxCoins(nums []int) int {
	n := len(nums)
	nums = append([]int{1}, append(nums, 1)...)
	dp := make([][]int, n+2)

	for i := range dp {
		dp[i] = make([]int, n+2)
		dp[i][i] = nums[i]
	}

	for length := 1; length < n+1; length++ {
		for start := 1; start < n-length+2; start++ {
			end := start + length - 1 // n-length+1
			for k := start; k < end+1; k++ {
				dp[start][end] = max2(dp[start][end], dp[start][k-1]+dp[k+1][end]+nums[k]*nums[start-1]*nums[end+1])
			}
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

// brute force, time limit exceeded
// func maxCoins(nums []int) int {
// 	curr, score := 0, 0
// 	solver(nums, curr, &score)

// 	return score
// }

// func solver(nums []int, curr int, score *int) {
// 	if len(nums) == 0 && curr > *score {
// 		*score = curr
// 		return
// 	}

// 	for i := range nums {
// 		tmp := nums[i]
// 		if i != 0 {
// 			tmp *= nums[i-1]
// 		}
// 		if i != len(nums)-1 {
// 			tmp *= nums[i+1]
// 		}
// 		solver(append(append([]int{}, nums[:i]...), nums[i+1:]...), curr+tmp, score)
// 	}
// }
