package main

import "fmt"

func main() {
	s := "aaaa"
	fmt.Println(partition(s))
}

func partition(s string) [][]string {
	if len(s) == 0 {
		return [][]string{}
	}
	if len(s) == 1 {
		return [][]string{[]string{s}}
	}
	ans := make([][]string, 0)
	solver(&ans, []string{}, s)
	return ans
}

func solver(ans *[][]string, prev []string, remain string) {
	if len(remain) == 0 {
		*ans = append(*ans, append([]string{}, prev...))
	}

	for i := 0; i < len(remain); i++ {
		if IsPalindrome(remain[0 : i+1]) {
			solver(ans, append(prev, remain[0:i+1]), remain[i+1:])
		}
	}
}

func IsPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
