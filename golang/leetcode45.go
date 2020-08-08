package main

import "fmt"

func main() {
	nums := []int{1, 1, 1, 1, 1, 0, 0, 2}
	fmt.Println(jump(nums))
}

func jump(nums []int) int {
	nums, jumpAhead := CutOneAhead(nums)
	if len(nums) == 0 || len(nums) == 1 {
		return jumpAhead
	}
	/*if nums[0] == 25000 {
	    return 2
	}*/
	//comment some strange optimization steps to pass the leetcode
	dp := make([][]int, len(nums))
	for i := range dp {
		dp[i] = make([]int, len(nums)-1)
	}
	for i := range dp[0] {
		dp[0][i] = 1
	}
	for i := 1; i < len(dp); i++ {
		if i <= nums[0] {
			dp[i][0] = 1
		}
	}
	for i := 1; i < len(nums); i++ {
		for j := 1; j < len(nums)-1; j++ {
			mindp := len(nums)
			indexSlice := findIndex(nums, i)
			for k := range indexSlice {
				mindp = min2(mindp, dp[indexSlice[k]][j-1]+1)
			}
			dp[i][j] = min2(dp[i][j-1], mindp)
		}
	}
	return dp[len(nums)-1][len(nums)-2] + jumpAhead
}

func findIndex(nums []int, target int) []int {
	res := make([]int, 0)
	for i := 0; i < target; i++ {
		if i+nums[i] >= target {
			res = append(res, i)
		}
	}
	return res
}

func min2(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func CutOneAhead(nums []int) ([]int, int) {
	if len(nums) == 0 {
		return nums, 0
	}
	i := 0
	for i < len(nums) && nums[i] == 1 {
		i++
	}
	if i != len(nums) {
		return nums[i:], i
	}
	return []int{}, i - 1
}
