/*
 * @lc app=leetcode id=376 lang=golang
 *
 * [376] Wiggle Subsequence
 */

// @lc code=start
func wiggleMaxLength(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	trend := 0 // if trend == 1, go up, if trend == -1 go down
	prevNum := nums[0]
	length := 1

	for i := 1; i < len(nums); i++ {
		if trend == 0 {
			if nums[i] > prevNum {
				trend = 1
				length++
			} else if nums[i] < prevNum {
				trend = -1
				length++
			}
			prevNum = nums[i]
			continue
		}
		if trend == 1 {
			if nums[i] < prevNum {
				trend = -1
				length++
			}
			prevNum = nums[i]
			continue
		}
		if trend == -1 {
			if nums[i] > prevNum {
				trend = 1
				length++
			}
			prevNum = nums[i]
			continue
		}
	}
	return length
}

// @lc code=end
