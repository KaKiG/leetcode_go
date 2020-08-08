package main

import "fmt"

func main() {
	s := "ca" // "ab"
	t := "ab" // "ca"
	fmt.Println(isIsomorphic(s, t))
}

// one pass (actually more) solution, faster
// optimize some if else conditions
func isIsomorphic(s string, t string) bool {
	charDic := make(map[byte]byte)

	for i := range s {
		_, exist := charDic[s[i]]
		if !exist {
			for _, v := range charDic {
				if v == t[i] {
					return false
				}
			}
			charDic[s[i]] = t[i]
		} else {
			if charDic[s[i]] != t[i] {
				return false
			}
		}
	}
	return true
}

// naive solution
// func isIsomorphic(s string, t string) bool {
// 	return solver(s, t) && solver(t, s)
// }

// func solver(s, t string) bool {
// 	charDic := make(map[string]string)

// 	for i := range s {
// 		_, exist := charDic[s[i:i+1]]
// 		if !exist {
// 			charDic[s[i:i+1]] = t[i : i+1]
// 		} else {
// 			if charDic[s[i:i+1]] != t[i:i+1] {
// 				return false
// 			}
// 		}
// 	}
// 	return true
// }
