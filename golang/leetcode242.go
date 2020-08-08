package main

import "fmt"

func main() {
	s := "rat"
	t := "car"

	fmt.Println(isAnagram(s, t))
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	letterDic := make(map[byte]int)
	for i := range s {
		letterDic[s[i]]++
		letterDic[t[i]]--
	}

	for _, v := range letterDic {
		if v != 0 {
			return false
		}
	}
	return true
}
