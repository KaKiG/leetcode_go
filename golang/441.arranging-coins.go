/*
 * @lc app=leetcode id=441 lang=golang
 *
 * [441] Arranging Coins
 */

// @lc code=start
func arrangeCoins(n int) int {
	left := 0
	right := n
	mid := 0

	if n == 1 {
		return 1
	}

	for left < right {
		mid = (left + right) / 2
		if (2*n < (1+mid)*mid) && (2*n >= mid*(mid-1)) {
			return mid - 1
		}
		if (2*n >= (1+mid)*mid) && (2*n < (2+mid)*(mid+1)) {
			return mid
		}
		if 2*n > (1+mid)*mid {
			left = mid
		} else {
			right = mid - 1
		}
	}
	return mid
}

// @lc code=end
