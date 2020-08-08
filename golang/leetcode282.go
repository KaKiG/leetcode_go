package main

import (
	"fmt"
	"strconv"
)

func main() {
	num := "105"
	target := 5

	fmt.Println(addOperators(num, target))
}

func addOperators(num string, target int) []string {
	ans := make([]string, 0)
	solver(num, "", 0, 0, target, &ans)
	return ans
}

// a recursive solver
func solver(num string, currExp string, prev, curr int, target int, ans *[]string) {
	if len(num) == 0 {
		if curr == target {
			*ans = append(*ans, currExp)
		}
		return
	}

	// dealing with only one number each time
	for i := 1; i < len(num)+1; i++ {
		// numbers like 0121 is not allowed
		if i > 1 && num[0] == '0' {
			break
		}
		val, _ := strconv.Atoi(num[:i])
		if len(currExp) == 0 {
			solver(num[i:], num[:i], val, val, target, ans)
		} else {
			// here we preserve prev value, because we when encounter *, it should be calculated first
			solver(num[i:], currExp+"+"+num[:i], val, curr+val, target, ans)
			solver(num[i:], currExp+"-"+num[:i], -val, curr-val, target, ans)
			solver(num[i:], currExp+"*"+num[:i], prev*val, curr-prev+prev*val, target, ans)
		}
	}
}
