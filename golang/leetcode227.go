package main

import "fmt"

func main() {
	s := "2*3+4"
	fmt.Println(calculate(s))
}

// sign = 1 is '+', sign = -1 is '-', sign = 3 is '*', sign = 4 is '4'
func calculate(s string) int {
	num := 0
	stack := make([]int, 0)
	for i := 0; i < len(s); i++ {

		switch s[i] {
		case '1', '2', '3', '4', '5', '6', '7', '8', '9', '0':
			j := i
			num = 0
			for j < len(s) && int(s[j]) >= 48 && int(s[j]) <= 57 {
				num = num*10 + int(s[j]-'0')
				j++
			}
			i = j - 1
			if len(stack) == 0 {
				stack = append(stack, num)
				continue
			}

			if stack[len(stack)-1] == 3 || stack[len(stack)-1] == 4 {
				prevNum := stack[len(stack)-2]
				res := 0
				if stack[len(stack)-1] == 3 {
					res = prevNum * num
				} else {
					res = prevNum / num
				}
				stack = stack[:len(stack)-2]
				stack = append(stack, res)
			} else {
				stack = append(stack, num)
			}
		case '+':
			stack = append(stack, 1)
		case '-':
			stack = append(stack, -1)
		case '*':
			stack = append(stack, 3)
		case '/':
			stack = append(stack, 4)
		}
	}
	res := stack[0]
	for i := 1; i < len(stack); i += 2 {
		res += stack[i] * stack[i+1]
	}
	return res

}
