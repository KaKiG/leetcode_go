package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := "3[a]2[bc]"
	fmt.Println(decodeString(s))
	s2 := "3[a2[c]]"
	fmt.Println(decodeString(s2))
	s3 := "2[abc]3[cd]ef"
	fmt.Println(decodeString(s3))
	s4 := "abc3[cd]xyz"
	fmt.Println(decodeString(s4))
	s5 := "100[leetcode]"
	fmt.Println(decodeString(s5))
}

func decodeString(s string) string {
	bracket_ind := make([]int, 0)
	ret_str := ""

	for i := range s {
		if s[i] == '[' {
			ret_str += string(s[i])
			bracket_ind = append(bracket_ind, len(ret_str)-1)
		} else if s[i] == ']' {
			start := bracket_ind[len(bracket_ind)-1]
			bracket_ind = bracket_ind[:len(bracket_ind)-1]
			num_start := identifyNumbers(ret_str, start-1)
			tmp_str := ret_str[num_start:]
			ret_str = ret_str[:num_start]
			ret_str += decodeStringSingle(tmp_str)
		} else {
			ret_str += string(s[i])
		}

	}

	return ret_str
}

func decodeStringSingle(s string) string {
	num := 0
	i := 0
	for i = range s {
		if s[i] == '[' {
			num, _ = strconv.Atoi(s[0:i])
			break
		}
	}
	ret_str := ""
	for j := 0; j < num; j++ {
		ret_str += s[i+1:]
	}
	return ret_str
}

func identifyNumbers(s string, start int) int {
	for i := start; i >= 0; i-- {
		if s[i] < '0' || s[i] > '9' {
			return i + 1
		}
	}
	return 0
}
