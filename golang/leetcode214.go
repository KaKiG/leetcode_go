package main

import "fmt"

func main() {
	s := "abac"
	fmt.Println(shortestPalindrome(s))
}

// kmp can be used
// T = s + $ + reverse(s)
// compute memo value
// memo[T[len(T)-1]] is the longest palindrome begins at s[0]
// return the result

func shortestPalindrome(s string) string {
	if len(s) == 0 || len(s) == 1 {
		return s
	}

	new := s + "$" + reverseStr(s)

	memoNew := memo(new)
	length := memoNew[len(memoNew)-1]
	return reverseStr(s[length:len(s)]) + s
}

// func shortestPalindrome(s string) string {
// 	if len(s) == 0 || len(s) == 1 {
// 		return s
// 	}

// 	for i := len(s); i > 0; i-- {
// 		if isPalindrome(s[0:i]) {
// 			return reverseStr(s[i:len(s)]) + s
// 		}
// 	}
// 	return reverseStr(s[1:len(s)]) + s
// }

func memo(s string) []int {
	memo := make([]int, len(s)+1)
	j := 0
	for i := 2; i <= len(s); i++ {
		for j > 0 && s[i-1] != s[j] {
			j = memo[j]
		}
		if s[i-1] == s[j] {
			j++
		}
		memo[i] = j
	}
	return memo[1:]
}

// func isPalindrome(s string) bool {
// 	right := len(s) - 1
// 	left := 0
// 	for left < right {
// 		if s[right] != s[left] {
// 			return false
// 		}
// 		right--
// 		left++
// 	}
// 	return true
// }

func reverseStr(s string) string {
	res := ""
	for i := range s {
		res = s[i:i+1] + res
	}
	return res
}
