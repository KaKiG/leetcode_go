package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 1}
	fmt.Println(rob(nums))
}

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	if len(nums) == 2 {
		return max2(nums[0], nums[1])
	}

	return max2(rob1(nums[1:]), rob1(nums[:len(nums)-1]))

}

func rob1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	if len(nums) == 2 {
		return max2(nums[0], nums[1])
	}

	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = max2(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		dp[i] = max2(dp[i-1], dp[i-2]+nums[i])
	}

	return dp[len(nums)-1]

}

func max2(a, b int) int {
	if a > b {
		return a
	}
	return b
}
