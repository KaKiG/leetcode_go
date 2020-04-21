package main

import (
	"fmt"
	"strings"
)

func main() {
	pattern := "abba"
	word := "dog dog dog dog"
	fmt.Println(wordPattern(pattern, word))
}

func wordPattern(pattern string, str string) bool {
	if len(pattern) == 0 && len(str) == 0 {
		return true
	}

	byteWord := make(map[byte]string)
	wordByte := make(map[string]byte)

	tmp := strings.Split(str, " ")

	if len(pattern) != len(tmp) {
		return false
	}

	for i := range pattern {
		_, exist1 := byteWord[pattern[i]]
		_, exist2 := wordByte[tmp[i]]

		if !exist1 && !exist2 {
			byteWord[pattern[i]] = tmp[i]
			wordByte[tmp[i]] = pattern[i]
		} else if !(exist1 && exist2 && byteWord[pattern[i]] == tmp[i] && wordByte[tmp[i]] == pattern[i]) {
			return false
		}
	}

	return true
}
