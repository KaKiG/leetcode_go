package main

func main() {

}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	if len(wordList) == 0 {
		return 0
	}

	wordMap := make(map[string]bool)

	for i := range wordList {
		wordMap[wordList[i]] = true
	}

	if !wordMap[endWord] {
		return 0
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
					return len(path) + 1
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
			for k := range levelMap {
				delete(wordMap, k)
			}
			levelMap = make(map[string]bool, 0)
			lenQ = len(queue)
		}
	}
	return 0
}
