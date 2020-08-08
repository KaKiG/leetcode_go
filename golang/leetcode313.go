package main

import "fmt"

func main() {
	n := 12
	primes := []int{2, 7, 13, 19}
	fmt.Println(nthSuperUglyNumber(n, primes))
}

func nthSuperUglyNumber(n int, primes []int) int {
	dp := make([]int, n)
	posn := make([]int, len(primes))
	dp[0] = 1

	for i := 1; i < len(dp); i++ {
		increase := []int{}

		dp[i], increase = minMultiplication(dp, posn, primes)

		for j := range increase {
			posn[increase[j]]++
		}
	}
	return dp[n-1]
}

func minMultiplication(dp, posn, primes []int) (int, []int) {
	res := make([]int, 0)
	minNum := dp[posn[0]] * primes[0]
	for i := range posn {
		val := dp[posn[i]] * primes[i]
		if val == minNum {
			res = append(res, i)
		}
		if val < minNum {
			minNum = val
			res = []int{i}
		}
	}
	return minNum, res
}
