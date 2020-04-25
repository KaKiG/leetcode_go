package main

import "fmt"

func main() {
	words := []string{"abcw", "baz", "foo", "bar", "xtfn", "abcdef"}
	fmt.Println(maxProduct(words))
}

// another way is to use bit operation
// for example for ba, start with 1, and do 1 << char - 'a', then do or expression, which is to keep the record
// then record for this word is  00000011 (assuming using 8 bit)
// to find overlapping between two words, we can basically use and operation between two recodes
func maxProduct(words []string) int {
	maxProd := 0
	record := make([]int, len(words))
	for i := range words {
		for j := range words[i] {
			record[i] = record[i] | 1<<(words[i][j]-'a')
		}
	}

	for i := range words {
		for j := i + 1; j < len(words); j++ {
			if record[i]&record[j] == 0 && len(words[i])*len(words[j]) > maxProd {
				maxProd = len(words[i]) * len(words[j])
			}
		}
	}
	return maxProd
}

// too slow
// func maxProduct(words []string) int {
// 	maxProd := 0

// 	for i := range words {
// 		for j := i + 1; j < len(words); j++ {
// 			l1 := len(words[i])
// 			l2 := len(words[j])
// 			if maxProd > 0 && l1*l2 <= maxProd {
// 				continue
// 			}
// 			if isValid(words[i], words[j]) {
// 				if l1*l2 > maxProd {
// 					maxProd = l1 * l2
// 				}
// 			}
// 		}
// 	}
// 	return maxProd
// }

// func isValid(a, b string) bool {
// 	letters := make(map[byte]bool)

// 	for i := range a {
// 		letters[a[i]] = true
// 	}

// 	for i := range b {
// 		if letters[b[i]] == true {
// 			return false
// 		}
// 	}
// 	return true
// }
