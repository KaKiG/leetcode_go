/*
 * @lc app=leetcode id=476 lang=golang
 *
 * [476] Number Complement
 */

// @lc code=start
func findComplement(num int) int {
	i := 0
	tmp := num
	hiBitMask := -1 << 31
	for i = 31; i >= 0; i-- {
		if (tmp & hiBitMask) == 0 {
			tmp <<= 1
		} else {
			break
		}
	}
	num = ^num
	if i == 31 {
		return num
	}
	num = num & (^(hiBitMask >> (30 - i)))
	return num
}

// @lc code=end
