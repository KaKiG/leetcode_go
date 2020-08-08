/*
 * @lc app=leetcode id=461 lang=golang
 *
 * [461] Hamming Distance
 */

// @lc code=start
func hammingDistance(x int, y int) int {
	count := 0
	x = x ^ y
	for i := 0; i < 31; i++ {
		count += (x & 1)
		x = x >> 1
	}
	return count
}

// @lc code=end
