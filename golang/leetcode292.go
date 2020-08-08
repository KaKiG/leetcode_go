package main

import "fmt"

func main() {
	n := 20
	fmt.Println(canWinNim(n))
}

func canWinNim(n int) bool {
	return n%4 != 0
}

// func canWinNim(n int) []bool {
// 	if n == 1 || n == 2 || n == 3 {
// 		return []bool{true}
// 	}

// 	dp := make([]bool, n)
// 	for i := 0; i < 3; i++ {
// 		dp[i] = true
// 	}

// 	for i := 3; i < n; i++ {
// 		dp[i] = !(dp[i-1] && dp[i-2] && dp[i-3])
// 	}

// 	return dp
// }
