package main

import "fmt"

func main() {
	s := "226"
	fmt.Println(numDecodings(s))
}

func numDecodings(s string) int {
	if len(s) == 0 {
		return 0
	}
	dp := make([]int, len(s))
	if s[0:1] == "0" {
		return 0
	}
	dp[0] = 1
	for i := 1; i < len(s); i++ {
		if s[i:i+1] == "0" {
			if !(s[i-1:i] == "1" || s[i-1:i] == "2") {
				return 0
			}
			if i < 2 {
				dp[i] = 1
			} else {
				dp[i] = dp[i-2]
			}
		} else {
			if (s[i:i+1] > "0" && s[i:i+1] <= "9" && s[i-1:i] == "1") || (s[i:i+1] > "0" && s[i:i+1] <= "6" && s[i-1:i] == "2") {
				if i >= 2 {
					dp[i] = dp[i-1] + dp[i-2]
				} else {
					dp[i] = dp[i-1] + 1
				}
			} else {
				dp[i] = dp[i-1]
			}
		}
	}
	fmt.Println(dp)
	return dp[len(s)-1]
}
