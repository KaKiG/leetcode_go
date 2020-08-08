import "strconv"

/*
 * @lc app=leetcode id=443 lang=golang
 *
 * [443] String Compression
 */

// @lc code=start
func compress(chars []byte) int {
	if len(chars) == 0 {
		return 0
	}
	currLetter := chars[0]
	pos := 0
	for i := 0; i < len(chars); i++ {
		count := 1
		for i < len(chars) && chars[i] == currLetter {
			i++
			count++
		}
		pos++
		chars[pos] = strconv.Itoa(count)
		pos++
		currLetter = chars[i]
	}
	return pos
}

// @lc code=end
