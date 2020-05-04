package main

import (
	"fmt"
	"math"
)

func main() {
	n := 12
	fmt.Println(numSquares(n))
}

func numSquares(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	dp := make([]int, n)
	dp[0] = 1

	for i := 1; i < n; i++ {
		factor := int(math.Sqrt(float64(i + 1)))
		if i+1 == factor*factor {
			dp[i] = 1
			continue
		}
		minNum := n
		for j := 1; j <= factor; j++ {
			if dp[i-j*j] < minNum {
				minNum = dp[i-j*j]
			}
		}
		dp[i] = minNum + 1
	}
	return dp[n-1]
}
