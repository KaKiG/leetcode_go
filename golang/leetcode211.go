package main

func main() {

}

type WordDictionary struct {
	wordMap map[int][]string
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
	var obj WordDictionary
	obj.wordMap = make(map[int][]string)
	return obj
}

/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string) {
	if len(word) != 0 {
		_, exist := this.wordMap[len(word)]
		if exist {
			this.wordMap[len(word)] = append(this.wordMap[len(word)], word)
		} else {
			this.wordMap[len(word)] = []string{word}
		}
	}
}

/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {
	if len(this.wordMap[len(word)]) == 0 {
		return false
	}

	for _, v := range this.wordMap[len(word)] {
		if compare(v, word) {
			return true
		}
	}
	return false
}

func compare(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	for i := range s {
		if s[i] != t[i] && !(s[i] == '.' || t[i] == '.') {
			return false
		}
	}
	return true
}

// type WordDictionary struct {
// 	wordStr []string
// }

// // too slow
// /** Initialize your data structure here. */
// func Constructor() WordDictionary {
// 	var obj WordDictionary
// 	return obj
// }

// /** Adds a word into the data structure. */
// func (this *WordDictionary) AddWord(word string) {
// 	this.wordStr = append(this.wordStr, word)
// }

// /** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
// func (this *WordDictionary) Search(word string) bool {
// 	for _, v := range this.wordStr {
// 		if compare(v, word) {
// 			return true
// 		}
// 	}
// 	return false
// }

// func compare(s, t string) bool {
// 	if len(s) != len(t) {
// 		return false
// 	}

// 	for i := range s {
// 		if s[i] != t[i] && !(s[i] == '.' || t[i] == '.') {
// 			return false
// 		}
// 	}
// 	return true
// }

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
