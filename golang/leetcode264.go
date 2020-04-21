package main

import "fmt"

func main() {
	n := 1690
	fmt.Println(nthUglyNumber(n))
}

// For this dp solution, we can imagine that, we want to find the ith ugly number
// it is smallest number which is the 2*, 3* or 5* by some preivous ugly number but greater than (i-1)th ugly number
// each time we use 2,3 or 5, we should move corresponding dp index by one, because that is the minimum
// also there should be no *continue* in each *if* statement
// because if one number is both two minimal, we should move both index by one
func nthUglyNumber(n int) int {
	if n == 1 {
		return 1
	}
	res := make([]int, n)
	res[0] = 1
	two, three, five := 0, 0, 0
	for i := 1; i < n; i++ {
		res[i] = min2(res[two]*2, min2(res[three]*3, res[five]*5))
		if res[i] == res[two]*2 {
			two++
		}
		if res[i] == res[three]*3 {
			three++
		}
		if res[i] == res[five]*5 {
			five++
		}
	}
	return res[n-1]
}

func min2(a, b int) int {
	if a < b {
		return a
	}
	return b
}
