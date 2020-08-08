package main

import "fmt"

func main() {
	s := "loveleetcode"
	fmt.Println(firstUniqChar(s))
}

func firstUniqChar(s string) int {
	count := make([]int, 26)

	for i := range s {
		if count[s[i]-'a'] == 0 {
			count[s[i]-'a'] = i + 1
		} else {
			count[s[i]-'a'] = -1
		}
	}
	minIdx := len(s) + 1
	for i := range count {
		if count[i] != -1 && count[i] != 0 && count[i] < minIdx {
			minIdx = count[i]
		}
	}
	if minIdx == len(s)+1 {
		return -1
	}
	return minIdx - 1
}
