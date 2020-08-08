package main

import "fmt"

func main() {
	s := "0P"
	fmt.Println(preStr(s))
	fmt.Println(isPalindrome(s))
}

func isPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}
	s = preStr(s)
	left := 0
	right := len(s) - 1
	for left <= right {
		if s[left] == s[right] {
			left++
			right--
			continue
		}
		return false
	}
	return true
}

func preStr(s string) string {
	postStr := ""
	for i := range s {
		if s[i] <= 122 && s[i] >= 97 || s[i] <= 57 && s[i] >= 48 {
			postStr += string(s[i])
			continue
		}
		if s[i] >= 65 && s[i] <= 90 {
			postStr += string(s[i] + 32)
		}
	}
	return postStr
}
