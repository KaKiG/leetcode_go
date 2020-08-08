/*
 * @lc app=leetcode id=377 lang=golang
 *
 * [377] Combination Sum IV
 */

// @lc code=start
// brute force solution
// func combinationSum4(nums []int, target int) int {
// 	ans := 0
// 	solver(target, nums, &ans)
// 	return ans
// }

// func solver(target int, nums []int,  ans *int) {
// 	if target == 0 {
// 		*ans += 1
// 		return
// 	}
// 	if target < 0 {
// 		return
// 	}

// 	for i := range nums {
// 		solver(target-nums[i], nums, ans)
// 	}
// }

// dp solution
func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)

	for i := 1; i <= target; i++ {
		for j := range nums {
			if i-nums[j] > 0 {
				dp[i] += dp[i-nums[j]]
				continue
			}
			if i == nums[j] {
				dp[i]++
			}
		}
	}
	return dp[target]
}

// @lc code=end
