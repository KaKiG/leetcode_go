/*
 * @lc app=leetcode id=500 lang=golang
 *
 * [500] Keyboard Row
 */

// @lc code=start
func findWords(words []string) []string {
	res := make([]string, 0)
	row1Map := map[byte]bool{
		'q': true,
		'Q': true,
		'w': true,
		'W': true,
		'e': true,
		'E': true,
		'r': true,
		'R': true,
		't': true,
		'T': true,
		'y': true,
		'Y': true,
		'u': true,
		'U': true,
		'i': true,
		'I': true,
		'o': true,
		'O': true,
		'p': true,
		'P': true,
	}
	row2Map := map[byte]bool{
		'a': true,
		'A': true,
		's': true,
		'S': true,
		'd': true,
		'D': true,
		'f': true,
		'F': true,
		'g': true,
		'G': true,
		'h': true,
		'H': true,
		'j': true,
		'J': true,
		'k': true,
		'K': true,
		'l': true,
		'L': true,
	}

	for i := range words {
		row1 := false
		row2 := false
		row3 := false
		sameRow := true
		for j := range words[i] {
			if row1Map[words[i][j]] {
				row1 = true
				if row2 || row3 {
					sameRow = false
					break
				}
			} else if row2Map[words[i][j]] {
				row2 = true
				if row1 || row3 {
					sameRow = false
					break
				}
			} else {
				row3 = true
				if row1 || row2 {
					sameRow = false
					break
				}
			}
		}
		if sameRow {
			res = append(res, words[i])
		}
	}
	return res
}

// @lc code=end
