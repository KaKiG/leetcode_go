/*
 * @lc app=leetcode id=287 lang=golang
 *
 * [287] Find the Duplicate Number
 */

// @lc code=start
func findDuplicate(nums []int) int {
	slow := nums[0]
	fast := nums[nums[0]]
	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}

	// draw a graph to show that this move could reach entry
	fast = 0
	for fast != slow {
		fast = nums[fast]
		slow = nums[slow]
	}

	return fast
}

// @lc code=end
