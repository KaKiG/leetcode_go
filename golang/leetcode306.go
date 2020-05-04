package main

import (
	"fmt"
	"strconv"
)

func main() {
	num := "112358"
	fmt.Println(isAdditiveNumber(num))
}

func isAdditiveNumber(num string) bool {
	ans := false
	solver(num, []int{}, &ans)
	return ans
}

func solver(num string, curr []int, ans *bool) {
	if len(num) == 0 && len(curr) > 2 {
		*ans = true
		return
	}

	for i := 1; i < len(num)+1; i++ {
		if i > 1 && num[0] == '0' {
			break
		}
		val, _ := strconv.Atoi(num[:i])
		if len(curr) < 2 {
			solver(num[i:], append(curr, val), ans)
		} else {
			if val == curr[len(curr)-1]+curr[len(curr)-2] {
				solver(num[i:], append(curr, val), ans)
			}
		}
	}
}
