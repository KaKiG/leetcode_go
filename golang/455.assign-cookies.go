import "sort"

/*
 * @lc app=leetcode id=455 lang=golang
 *
 * [455] Assign Cookies
 */

// @lc code=start
func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	g_iter := 0
	s_iter := 0
	count := 0
	for g_iter != len(g) && s_iter != len(s) {
		if s[s_iter] >= g[g_iter] {
			count++
			s_iter++
			g_iter++
		} else {
			s_iter++
		}
	}
	return count
}

// @lc code=end
