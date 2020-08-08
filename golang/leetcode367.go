package main

import "fmt"

func main() {
	num := 9
	fmt.Println(isPerfectSquare(num))
}

func isPerfectSquare(num int) bool {
	if num == 0 || num == 1 {
		return true
	}

	left := 1
	right := num / 2
	mid := 0
	for left < right {
		mid = (left + right) / 2
		tmp := mid * mid
		if tmp == num {
			return true
		}

		if tmp < num {
			left = mid + 1
		}

		if tmp > num {
			right = mid
		}
	}

	if right*right == num {
		return true
	}
	return false
}
