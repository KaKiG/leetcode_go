package main

/*
 * @lc app=leetcode id=459 lang=golang
 *
 * [459] Repeated Substring Pattern
 */

// @lc code=start
// func main() {
// 	s := "abac"
// 	fmt.Println(repeatedSubstringPattern(s))
// }
func repeatedSubstringPattern(s string) bool {
	memoArr := memo(s)
	// fmt.Println(memoArr)
	return (memoArr[len(s)] != 0) && ((len(s) % (len(s) - memoArr[len(s)])) == 0)
}

// kmp solution to compute memo table
func memo(s string) []int {
	n := len(s)
	memo := make([]int, len(s)+1)

	j := 0
	for i := 2; i <= n; i++ {
		for j > 0 && s[i-1] != s[j] {
			j = memo[j]
		}
		if s[i-1] == s[j] {
			j++
		}
		memo[i] = j
	}
	return memo
}

// @lc code=end
