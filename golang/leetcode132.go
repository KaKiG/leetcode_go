package main

import "fmt"

func main() {
	s := "ababababababababababababcbabababababababababababa"
	fmt.Println(minCut(s))
}

func minCut(s string) int {
	length := len(s)
	dp := make([][]bool, length)
	for k := range dp {
		dp[k] = make([]bool, length)
	}

	cut := make([]int, length)
	for j := 0; j < length; j++ {
		cut[j] = j
		for i := 0; i <= j; i++ {
			if s[i] == s[j] && (j-i <= 1 || dp[i+1][j-1]) {
				dp[i][j] = true
				if i > 0 {
					cut[j] = min(cut[j], cut[i-1]+1)
				} else {
					cut[j] = 0
				}
			}
		}
	}
	return cut[length-1]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
