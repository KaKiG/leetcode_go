package main

import (
	"fmt"
	"strconv"
)

func main() {
	input := "11"
	fmt.Println(diffWaysToCompute(input))
}

func diffWaysToCompute(input string) []int {
	index := make([]int, 0)
	for i := range input {
		if input[i] == '+' || input[i] == '-' || input[i] == '*' {
			index = append(index, i)
		}
	}
	if len(index) == 0 {
		val, _ := strconv.Atoi(input)
		return []int{val}
	}

	var left, right, ans []int

	for _, i := range index {
		left = diffWaysToCompute(input[:i])
		right = diffWaysToCompute(input[i+1:])
		for j := range left {
			for k := range right {
				switch input[i] {
				case '+':
					ans = append(ans, left[j]+right[k])
				case '-':
					ans = append(ans, left[j]-right[k])
				case '*':
					ans = append(ans, left[j]*right[k])
				}
			}
		}

	}

	return ans
}
