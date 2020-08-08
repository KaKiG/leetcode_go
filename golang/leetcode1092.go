package main

import "fmt"

func main() {
	str1 := "bbbaaaba"
	str2 := "bbababbb"

	fmt.Println(shortestCommonSupersequence(str1, str2))
}

func shortestCommonSupersequence(str1 string, str2 string) string {
	dp := make([][]int, len(str1)+1)
	for i := range dp {
		dp[i] = make([]int, len(str2)+1)
	}
	res := ""
	for i := range dp {
		dp[i][0] = i
	}
	for i := range dp[0] {
		dp[0][i] = i
	}

	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[0]); j++ {
			if str1[i-1] == str2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				if dp[i-1][j] > dp[i][j-1] {
					dp[i][j] = dp[i][j-1] + 1
				} else {
					dp[i][j] = dp[i-1][j] + 1
				}
			}
		}
	}
	i := len(str1)
	j := len(str2)

	for {
		if str1[i-1] == str2[j-1] {
			res = str1[i-1:i] + res
			i--
			j--
		} else if dp[i-1][j] > dp[i][j-1] {
			res = str2[j-1:j] + res
			j--
		} else {
			res = str1[i-1:i] + res
			i--
		}

		if i == 0 && j == 0 {
			break
		} else if i == 0 {
			res = str2[:j] + res
			break
		} else if j == 0 {
			res = str1[:i] + res
			break
		}
	}
	return res
}
