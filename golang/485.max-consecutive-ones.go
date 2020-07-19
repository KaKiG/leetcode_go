/*
 * @lc app=leetcode id=485 lang=golang
 *
 * [485] Max Consecutive Ones
 */

// @lc code=start
func findMaxConsecutiveOnes(nums []int) int {
	maxCount := 0
	for i := 0; i < len(nums); i++ {
		tmpCount := 0
		for i < len(nums) && nums[i] == 1 {
			tmpCount++
			i++
		}
		if tmpCount > maxCount {
			maxCount = tmpCount
		}
	}
	return maxCount
}

// @lc code=end
