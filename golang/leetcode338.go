package main

import "fmt"

func main() {
	num := 10
	fmt.Println(countBits(num))
}

func countBits(num int) []int {
	dp := make([]int, num+1)
	tmp := 1
	for i := 1; i <= num; i++ {
		if tmp*2 <= i {
			tmp *= 2
		}
		dp[i] = dp[i-tmp] + 1
	}
	return dp
}
