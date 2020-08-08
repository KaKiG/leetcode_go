package main

import "fmt"

func main() {
	obj := Constructor()
	obj.Insert("hello")
	fmt.Println(obj.Search("hell"))
	fmt.Println(obj.Search("helloa"))
	fmt.Println(obj.Search("hello"))
	fmt.Println(obj.StartsWith("hell"))
	fmt.Println(obj.StartsWith("helloa"))
	fmt.Println(obj.StartsWith("hello"))

}

type Trie struct {
	val  byte
	word []*List
}

type List struct {
	val  byte
	Next *List
}

/** Initialize your data structure here. */
func Constructor() Trie {
	var trie Trie
	trie.val = '$'
	trie.word = nil
	return trie
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	wordNode := make([]List, len(word))
	for i := range word {
		wordNode[i].val = word[i]
		if i != len(word)-1 {
			wordNode[i].Next = &wordNode[i+1]
		}
	}
	this.word = append(this.word, &wordNode[0])
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	for j := range this.word {
		if this.word[j].val == word[0] {
			curr := this.word[j]
			for i := range word {
				if curr.val == word[i] {
					curr = curr.Next
					if curr == nil {
						if i == len(word)-1 {
							return true
						} else {
							break
						}
					}
				} else {
					break
				}
			}
		}
	}
	return false
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	for j := range this.word {
		if this.word[j].val == prefix[0] {
			curr := this.word[j]
			for i := range prefix {
				if curr.val == prefix[i] {
					if i == len(prefix)-1 {
						return true
					}
					curr = curr.Next
					if curr == nil {
						break
					}
				} else {
					break
				}
			}
		}
	}
	return false
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
