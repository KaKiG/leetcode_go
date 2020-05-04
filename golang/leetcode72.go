package main

import "fmt"

func main() {
	word1 := "intention"
	word2 := "execution"
	fmt.Println(minDistance(word1, word2))
}

func minDistance(word1 string, word2 string) int {
	dp := make([][]int, len(word1)+1)
	for i := range dp {
		dp[i] = make([]int, len(word2)+1)
	}
	dp[0][0] = 0
	for i := 1; i < len(word1)+1; i++ {
		dp[i][0] = dp[i-1][0] + 1
	}
	for i := 1; i < len(word2)+1; i++ {
		dp[0][i] = dp[0][i-1] + 1
	}

	for i := 1; i < len(word1)+1; i++ {
		for j := 1; j < len(word2)+1; j++ {
			dp[i][j] = min3(1+dp[i-1][j], 1+dp[i][j-1], Cost(word1[i-1:i], word2[j-1:j])+dp[i-1][j-1])
		}
	}
	return dp[len(word1)][len(word2)]
}

func min2(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func min3(a, b, c int) int {
	return min2(min2(a, b), c)
}

func Cost(a, b string) int {
	if a == b {
		return 0
	}
	return 1
}
