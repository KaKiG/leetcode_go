package main

import "fmt"

func main() {
	// s := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaabaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	// wordDict := []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"}

	// s := "catsanddog"
	// wordDict := []string{"cat", "cats", "and", "sand", "dog"}

	s := "pineapplepenapple"
	wordDict := []string{"apple", "pen", "applepen", "pine", "pineapple"}
	str := wordBreak(s, wordDict)
	for i := range str {
		fmt.Println(str[i])
	}
}

func wordBreak(s string, wordDict []string) []string {
	wordMap := make(map[string]bool)
	for i := range wordDict {
		wordMap[wordDict[i]] = true
	}

	dp := make([]bool, len(s)+1)
	path := make([][]int, len(s)+1)
	for i := range dp {
		if wordMap[s[:i]] {
			dp[i] = true
			path[i] = append(path[i], i)
		}
	}

	for i := 1; i < len(dp); i++ {
		for j := 1; j < i; j++ {
			if dp[j] && wordMap[s[j:i]] {
				dp[i] = true
				path[i] = append(path[i], j)
			}
		}
	}
	return PartitionWord(path, len(s), s, map[int][]string{})
}

func PartitionWord(pos [][]int, i int, s string, strMap map[int][]string) []string {
	res := []string{}
	if _, exist := strMap[i]; exist {
		return strMap[i]
	}
	for j := 0; j < len(pos[i]); j++ {
		if len(pos[i]) != 0 && pos[i][j] == i {
			res = append(res, s[:i])
			continue
		}
		prev := PartitionWord(pos, pos[i][j], s, strMap)
		for k := range prev {
			res = append(res, prev[k]+" "+s[pos[i][j]:i])
		}
	}
	strMap[i] = res

	return res
}
