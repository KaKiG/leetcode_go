/*
 * @lc app=leetcode id=509 lang=golang
 *
 * [509] Fibonacci Number
 */

// @lc code=start
func fib(N int) int {
	if N == 0 || N == 1 {
		return N
	}
	arr := make([]int, N+1)
	arr[1] = 1

	for i := 2; i <= N; i++ {
		arr[i] = arr[i-1] + arr[i-2]
	}
	return arr[N]
}

// @lc code=end
