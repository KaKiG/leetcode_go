/*
 * @lc app=leetcode id=448 lang=golang
 *
 * [448] Find All Numbers Disappeared in an Array
 */

// @lc code=start
func findDisappearedNumbers(nums []int) []int {
	arr := make([]int, len(nums))
	for i := range nums {
		arr[nums[i]-1] = 1
	}
	j := 0
	for i := range arr {
		if arr[i] == 0 {
			nums[j] = i + 1
			j++
		}
	}
	return nums[:j]
}

// @lc code=end
