import "strconv"

/*
 * @lc app=leetcode id=412 lang=golang
 *
 * [412] Fizz Buzz
 */

// @lc code=start

func fizzBuzz(n int) []string {
	res := make([]string, n)
	for i := 1; i <= n; i++ {
		if isMultipleThree(i) && isMultipleFive(i) {
			res[i-1] = "FizzBuzz"
			continue
		}
		if isMultipleThree(i) {
			res[i-1] = "Fizz"
			continue
		}
		if isMultipleFive(i) {
			res[i-1] = "Buzz"
			continue
		}
		res[i-1] = strconv.Itoa(i)
	}
	return res
}

func isMultipleFive(n int) bool {
	if n%5 == 0 {
		return true
	}
	return false
}

func isMultipleThree(n int) bool {
	if n%3 == 0 {
		return true
	}
	return false
}

// @lc code=end
