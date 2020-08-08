import "strconv"

/*
 * @lc app=leetcode id=504 lang=golang
 *
 * [504] Base 7
 */

// @lc code=start
func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}
	s := ""
	for num != 0 {
		remainder := num % 7
		num /= 7
		if num == 0 {
			s = strconv.Itoa(remainder) + s
		} else {
			s = strconv.Itoa(abs(remainder)) + s
		}
	}

	return s
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// @lc code=end
