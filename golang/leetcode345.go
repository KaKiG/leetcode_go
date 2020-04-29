package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "leetcode"
	// s := "a a"
	fmt.Println(reverseVowels(s))
}

// could use trick r := []rune{s}
func reverseVowels(s string) string {
	if len(s) < 2 {
		return s
	}
	strSlice := strings.Split(s, "")
	left := 0
	right := len(s) - 1
	for left < right {
		if !isVowel(strSlice[left][0]) {
			left++
		}
		if !isVowel(strSlice[right][0]) {
			right--
		}
		if isVowel(strSlice[left][0]) && isVowel(strSlice[right][0]) {
			strSlice[left], strSlice[right] = strSlice[right], strSlice[left]
			left++
			right--
		}
	}
	return strings.Join(strSlice, "")
}

// func reverseVowels(s string) string {
// 	if len(s) < 2 {
// 		return s
// 	}
// 	leftEnd := 0
// 	rightEnd := len(s) - 1
// 	leftStart := leftEnd
// 	rightStart := rightEnd
// 	resLeft := ""
// 	resRight := ""
// 	for leftEnd < rightEnd {
// 		for !(isVowel(s[leftEnd]) && isVowel(s[rightEnd])) && leftEnd < rightEnd {
// 			if !isVowel(s[leftEnd]) {
// 				leftEnd++
// 			}
// 			if !isVowel(s[rightEnd]) {
// 				rightEnd--
// 			}
// 		}
// 		if leftEnd < rightEnd {
// 			resLeft = resLeft + s[leftStart:leftEnd] + string(s[rightEnd])
// 			resRight = string(s[leftEnd]) + s[rightEnd+1:rightStart+1] + resRight
// 			leftEnd++
// 			rightEnd--
// 			leftStart, rightStart = leftEnd, rightEnd
// 		}

// 		if leftEnd >= rightEnd {
// 			resLeft = resLeft + s[leftStart:leftEnd]
// 			resRight = s[rightEnd+1:rightStart+1] + resRight
// 			if leftEnd == rightEnd {
// 				resLeft += string(s[leftEnd])
// 			}
// 		}

// 	}
// 	return resLeft + resRight
// }

func isVowel(s byte) bool {
	switch s {
	case 'a':
		return true
	case 'e':
		return true
	case 'i':
		return true
	case 'o':
		return true
	case 'u':
		return true
	case 'A':
		return true
	case 'E':
		return true
	case 'I':
		return true
	case 'O':
		return true
	case 'U':
		return true
	}
	return false
}
