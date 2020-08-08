package main

import "fmt"

func main() {
	s1 := "bcabc"
	fmt.Println(removeDuplicateLetters(s1))
	s2 := "cbacdcbc"
	fmt.Println(removeDuplicateLetters(s2))
	s3 := "abcacb"
	fmt.Println(removeDuplicateLetters(s3))
	s4 := "bbcaac"
	fmt.Println(removeDuplicateLetters(s4))
}

func removeDuplicateLetters(s string) string {
	totalLetterNumDic := make(map[byte]int)
	currLetterDic := make(map[byte]bool)
	letterStack := ""
	for i := range s {
		totalLetterNumDic[s[i]]++
	}

	for i := range s {
		totalLetterNumDic[s[i]]--
		if !currLetterDic[s[i]] {
			for len(letterStack) != 0 {
				prevLetter := letterStack[len(letterStack)-1]
				if (len(letterStack) != 0) && prevLetter > s[i] && totalLetterNumDic[prevLetter] > 0 {
					currLetterDic[prevLetter] = false
					letterStack = letterStack[:len(letterStack)-1]
				} else {
					break
				}
			}
			letterStack += string(s[i])
			currLetterDic[s[i]] = true
		}
	}
	return letterStack
}
