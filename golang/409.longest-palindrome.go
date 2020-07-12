/*
 * @lc app=leetcode id=409 lang=golang
 *
 * [409] Longest Palindrome
 */

// @lc code=start
func longestPalindrome(s string) int {
	arr := make([]int, 52)
	res := 0
	odd := 0

	for i := range s {
		if s[i] >= 'a' && s[i] <= 'z' {
			arr[s[i]-97]++
		} else {
			arr[s[i]-39]++
		}
	}

	for i := range arr {
		res += arr[i]
		if arr[i]%2 != 0 {
			res -= 1
			odd = 1
		}
	}
	return res + odd
}

// @lc code=end
