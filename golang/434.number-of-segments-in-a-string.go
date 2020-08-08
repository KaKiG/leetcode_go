/*
 * @lc app=leetcode id=434 lang=golang
 *
 * [434] Number of Segments in a String
 */

// @lc code=start
func countSegments(s string) int {
	seg := false
	count := 0
	for i := range s {
		if s[i] != ' ' && !seg {
			seg = true
		}
		if s[i] == ' ' && seg {
			count++
			seg = false
		}
	}
	if seg {
		count++
	}
	return count
}

// @lc code=end
