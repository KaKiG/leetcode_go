import (
	"strconv"
)

/*
 * @lc app=leetcode id=93 lang=golang
 *
 * [93] Restore IP Addresses
 */

// @lc code=start
func restoreIpAddresses(s string) []string {
	if len(s) > 12 || len(s) < 4 {
		return []string{}
	}
	res := make([]string, 0)
	getIpAddress(s, 3, &res, "")
	return res
}

func getIpAddress(s string, count int, res *[]string, curr string) {
	if count == 0 {
		if isValid(s) {
			*res = append(*res, curr+(s))
			return
		}
	}

	count = count - 1
	for i := 1; i < 4; i++ {
		tmp := curr
		if i <= len(s) {
			if isValid(s[:i]) && len(s[i:]) <= 3*(count+1) {
				tmp += s[:i]
				tmp += "."
				getIpAddress(s[i:], count, res, tmp)
			}
		}

	}
}

func isValid(s string) bool {
	num, _ := strconv.Atoi(s)
	if num >= 0 && num <= 255 && len(s) > 0 {
		if len(s) > 1 && s[0] == '0' {
			return false
		}
		return true
	}
	return false
}

// @lc code=end
