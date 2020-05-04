package main

import (
	"fmt"
)

func main() {
	beginWord := "hit"
	endWord := "cog"
	wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}
	fmt.Println(findLadders(beginWord, endWord, wordList))
}

func findLadders(beginWord, endWord string, wordList []string) [][]string {
	res := make([][]string, 0)
	if len(wordList) == 0 {
		return res
	}

	wordMap := make(map[string]bool)

	for i := range wordList {
		wordMap[wordList[i]] = true
	}

	if !wordMap[endWord] {
		return res
	}

	levelMap := make(map[string]bool, 0)
	queue := make([][]string, 0)
	queue = append(queue, []string{beginWord})
	lenQ := 1

	for len(queue) > 0 {
		path := queue[0]
		queue = queue[1:]
		prevWord := path[len(path)-1]
		for i := range prevWord {
			for c := 'a'; c <= 'z'; c++ {
				nextWord := prevWord[:i] + string(c) + prevWord[i+1:]
				if nextWord == endWord {
					path = append(path, nextWord)
					res = append(res, path)
					continue
				}
				if wordMap[nextWord] {
					newPath := make([]string, len(path))
					copy(newPath, path)
					newPath = append(newPath, nextWord)
					queue = append(queue, newPath)
					levelMap[nextWord] = true
				}
			}
		}

		lenQ--

		if lenQ == 0 {
			if len(res) > 0 {
				return res
			}
			for k := range levelMap {
				delete(wordMap, k)
			}
			levelMap = make(map[string]bool, 0)
			lenQ = len(queue)
		}
	}
	return res
}

//DFS is somehow slower in this problem.
// func findLadders(beginWord string, endWord string, wordList []string) [][]string {
// 	ans := make([][]string, 0)
// 	min := len(wordList) + 1
// 	solver(beginWord, endWord, wordList, []string{beginWord}, &ans, &min)
// 	return ans
// }
//
// func solver(begin, end string, list []string, curr []string, ans *[][]string, min *int) {
// 	if begin == end {
// 		if len(curr) < *min {
// 			*min = len(curr)
// 			*ans = (*ans)[:0]
// 			*ans = append(*ans, append([]string{}, curr...))
// 		} else if len(curr) == *min {
// 			*ans = append(*ans, append([]string{}, curr...))
// 		}
// 	}
// 	for i := range list {
// 		if changeOne(begin, list[i]) {
// 			solver(list[i], end, append(append([]string{}, list[:i]...), list[i+1:]...), append(curr, list[i]), ans, min)
// 		}
// 	}
// }
//
// func changeOne(s, t string) bool {
// 	count := 0
// 	for i := range s {
// 		if s[i] != t[i] {
// 			count++
// 		}
// 	}
// 	if count == 1 {
// 		return true
// 	}
// 	return false
// }
