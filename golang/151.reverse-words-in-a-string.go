package main

/*
 * @lc app=leetcode id=151 lang=golang
 *
 * [151] Reverse Words in a String
 */

// @lc code=start
// func main() {
// 	s := " s   hello word! "
// 	fmt.Println("[" + reverseWords(s) + "]")
// 	s1 := "  "
// 	fmt.Println("[" + reverseWords(s1) + "]")
// }
func reverseWords(s string) string {
	if len(s) == 0 {
		return s
	}
	res := ""
	for i := len(s) - 1; i >= 0; {
		word, ind := findNext(s, i)
		res += word
		if len(word) > 0 {
			res += " "
		}
		i = ind
	}
	if len(res) > 0 {
		return res[:len(res)-1]
	}
	return res
}

func findNext(s string, start int) (string, int) {
	i := start
	wordStart := start
	wordEnd := start
	for i >= 0 && s[i] == ' ' {
		i--
	}
	wordStart = i
	for i >= 0 && s[i] != ' ' {
		i--
	}
	wordEnd = i
	return s[wordEnd+1 : wordStart+1], i
}

// @lc code=end
