package main

import "fmt"

func main() {
	s := "Z"
	fmt.Println(titleToNumber(s))
}

func titleToNumber(s string) int {
	total := 0
	n := len(s)
	for i := range s {
		total += (int(s[i]) - 64) * pow26(n-i-1)
	}
	return total
}

func pow26(a int) int {
	res := 1
	for i := 0; i < a; i++ {
		res *= 26
	}
	return res
}
