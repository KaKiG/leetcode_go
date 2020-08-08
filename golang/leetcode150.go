package main

import (
	"fmt"
	"strconv"
)

func main() {
	tokens := []string{"4", "13", "5", "/", "+"}
	fmt.Println(evalRPN(tokens))
}

func evalRPN(tokens []string) int {
	stack := make([]int, 0)
	for i := range tokens {
		if tokens[i] == "+" {
			tmp1 := stack[len(stack)-2]
			tmp2 := stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, tmp1+tmp2)
		} else if tokens[i] == "-" {
			tmp1 := stack[len(stack)-2]
			tmp2 := stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, tmp1-tmp2)
		} else if tokens[i] == "*" {
			tmp1 := stack[len(stack)-2]
			tmp2 := stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, tmp1*tmp2)
		} else if tokens[i] == "/" {
			tmp1 := stack[len(stack)-2]
			tmp2 := stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, tmp1/tmp2)
		} else {
			v, _ := strconv.Atoi(tokens[i])
			stack = append(stack, v)
		}
	}
	return stack[0]
}
