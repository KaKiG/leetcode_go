import (
	"math"
)

/*
 * @lc app=leetcode id=507 lang=golang
 *
 * [507] Perfect Number
 */

// @lc code=start
func checkPerfectNumber(num int) bool {
	if num < 3 {
		return false
	}
	sum := 0
	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if (num % i) == 0 {
			sum += i
			sum += (num / i)
		}
	}
	return sum == (num - 1)
}

// @lc code=end
