/*
 * @lc app=leetcode id=520 lang=golang
 *
 * [520] Detect Capital
 */

// @lc code=start
func detectCapitalUse(word string) bool {
	allCap := true
	restUnCap := true

	for i := 0; i < len(word); i++ {
		if word[i] >= 'a' && word[i] <= 'z' {
			allCap = false
		} else {
			if i != 0 {
				restUnCap = false
			}
		}

		if !(allCap || restUnCap) {
			return false
		}
	}
	return true
}

// @lc code=end
