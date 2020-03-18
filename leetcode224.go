package main

import (
	"fmt"
	"strconv"
)

func main() {
	// s := "(1+(4+5+2)-3)+(6+8) "
	s := "   (7 - 8) - (   9 - 11)   "
	fmt.Println(calculate(s))
}

func calculate(s string) int {
	expre := ""
	indices := []int{}
	for i := range s {
		if s[i] == ' ' {
			i++
			continue
		}

		expre += s[i : i+1]

		if s[i] == '(' {
			indices = append(indices, len(expre)-1)
		}

		if s[i] == ')' {
			index := indices[len(indices)-1]
			indices = indices[:len(indices)-1]
			tmp := expre[index+1 : len(expre)-1]
			expre = expre[:index]
			val := strconv.Itoa(cal(tmp))
			expre += val
		}
	}
	return cal(expre)
}

func cal(s string) int {
	i := 0
	for i = 0; i < len(s); i++ {
		if i != 0 && (s[i] == '+' || s[i] == '-') {
			break
		}
	}
	tmp, _ := strconv.Atoi(s[0:i])
	res := tmp
	for ; i < len(s); i++ {
		if s[i] == ' ' {
			i++
			continue
		}
		if s[i] == '+' {
			j := i + 2
			val := 0
			for j = i + 2; j < len(s); j++ {
				if s[j] == '+' || s[j] == '-' {
					val, _ = strconv.Atoi(s[i+1 : j])
					break
				}
			}
			if j >= len(s)-1 {
				val, _ = strconv.Atoi(s[i+1 : len(s)])
			}
			res += val
			i = j - 1
		} else if s[i] == '-' {
			j := i + 2
			val := 0
			for j = i + 2; j < len(s); j++ {
				if s[j] == '+' || s[j] == '-' {
					val, _ = strconv.Atoi(s[i+1 : j])
					break
				}
			}
			if j >= len(s)-1 {
				val, _ = strconv.Atoi(s[i+1 : len(s)])
			}
			res -= val
			i = j - 1
		}
	}
	return res
}
